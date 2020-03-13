// UTILITIES
function heatMapColorforValue(value) {
    var h = (1.0 - value) * 240
    return "hsl(" + h + ", 100%, 50%)";
}

function fetchJSONFile(path, callback) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {
            var data = JSON.parse(httpRequest.responseText);
            if (callback) callback(data);
        }
    };
    httpRequest.open('GET', path);
    httpRequest.send();
}

function normalisePlace(place) {
    var lowerPlace = place.toLowerCase();

    // italian
    if (lowerPlace === 'valle d\'aosta/vallée d\'aoste') return 'aosta';
    else if (lowerPlace === 'bolzano/bozen') return 'bolzano';
    else if (lowerPlace === 'massa-carrara') return 'massa carrara';
    else if (lowerPlace === 'alto adige') return 'p.a. bolzano';
    else if (lowerPlace === 'trentino') return 'p.a. trento';

    // english
    if (lowerPlace === 'florence') return 'firenze';
    else if (lowerPlace === 'genoa') return 'genova';
    else if (lowerPlace === 'mantua') return 'mantova';
    else if (lowerPlace === 'milan') return 'milano';
    else if (lowerPlace === 'naples') return 'napoli';
    else if (lowerPlace === 'padua') return 'padova';
    else if (lowerPlace === 'rome') return 'roma';
    else if (lowerPlace === 'south sardinia') return 'sud sargegna';
    else if (lowerPlace === 'syracuse') return 'siracusa';
    else if (lowerPlace === 'turin') return 'torino';
    else if (lowerPlace === 'venice') return 'venezia';

    if (lowerPlace === 'latium') return 'lazio';
    else if (lowerPlace === 'apulia') return 'puglia';
    else if (lowerPlace === 'sardinia') return 'sardegna';
    else if (lowerPlace === 'sicily') return 'sicilia';
    else if (lowerPlace === 'south tyrol') return 'p.a. bolzano';
    else if (lowerPlace === 'south tyrol') return 'p.a. trento';

    return lowerPlace;
}

function paintMap(type, data, lastUpdate) {
    document.getElementById('lastUpdate').innerHTML = lastUpdate;
    var max = Math.max.apply(Math, Object.values(data));

    var endpoint = type === 'region'
                   ? '/limits_IT_regions.geojson'
                   : '/limits_IT_provinces.geojson';

    fetchJSONFile(endpoint, function(items) {
        if (geoLayer !== undefined) {
            geoLayer.remove();
        }

        geoLayer = L.geoJSON(items, {
           weight: 1,
           style: function(feature) {
               var place = type === 'region'
                           ? normalisePlace(feature.properties.reg_name)
                           : normalisePlace(feature.properties.prov_name);
               var cases = data[place];
               if (cases > 0) {
                   return {
                       color: heatMapColorforValue(Math.min(cases / max + 0.75, 1))
                   };
               }
           },
           onEachFeature: function(feature, layer) {
               var place = type === 'region'
                           ? normalisePlace(feature.properties.reg_name)
                           : normalisePlace(feature.properties.prov_name);
               var cases = data[place];
               var popupContent = '<strong>' + place + '</strong><br>';
               popupContent += 'Cases: ' + cases + '<br>';
               popupContent += '<small>Updated on ' + lastUpdate + '<small>';
               layer.bindPopup(popupContent);
           }
        })
        .addTo(map)
        .on('click', function(e) {
            document.getElementById('type-form').value = type;
            document.getElementById('value-form').value = type === 'region'
                    ? normalisePlace(e.sourceTarget.feature.properties.reg_name)
                    : normalisePlace(e.sourceTarget.feature.properties.prov_name);
            document.getElementById('search-form').submit();
        });
    });
}

// SEARCH FORM
document.getElementById('search-form').onsubmit = function(e) {
    document.getElementById('value-form').value = normalisePlace(document.getElementById('value-form').value);
};

// COVID DATA
history = {region: {}, province: {}};

// GEOMAP
var map = L.map('map').setView([41.8719, 12.5674], 6);

L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token=pk.eyJ1IjoiZmFiaW9jaWNlcmNoaWEiLCJhIjoiY2s3b2phNXhxMDlzbTNncDkzY3pkb2YxYSJ9.37ZdJEamTJMHvbticJ25CQ', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, ' +
        '<a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
        'Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    id: 'mapbox/light-v9'
}).addTo(map);

geoLayer = undefined;

Papa.parse('/dpc-covid19-ita-province.csv', {
    header: true,
    download: true,
    complete: function(results) {
        var date = new Date();
        var today = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
        var yesterday = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate() - 1).toString().padStart(2, '0');
        results.data.forEach(function(item) {
            var parsedDate = item.data.substr(0, 10);
            if (typeof history['province'][parsedDate] === 'undefined') history['province'][parsedDate] = {};
	    province = normalisePlace(item.denominazione_provincia);
            history['province'][parsedDate][province] = parseInt(item.totale_casi, 10);
        });
        var dataProvince = history['province'][today] || history['province'][yesterday];
        var lastUpdate = history['province'][today] ? today : yesterday;
        var sumCases = Object.values(dataProvince).reduce((a, b) => a + b, 0);
        document.getElementById('totalCases').innerHTML = sumCases.toLocaleString();

        paintMap('province', dataProvince, lastUpdate);
    }
});

Papa.parse('/dpc-covid19-ita-regioni.csv', {
    header: true,
    download: true,
    complete: function(results) {
        var date = new Date();
        var today = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
        var yesterday = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate() - 1).toString().padStart(2, '0');
        results.data.forEach(function(item) {
            var parsedDate = item.data.substr(0, 10);
            if (typeof history['region'][parsedDate] === 'undefined') history['region'][parsedDate] = {};
	    province = normalisePlace(item.denominazione_provincia);
            history['region'][parsedDate][province] = parseInt(item.totale_casi, 10);
        });
        var dataProvince = history['region'][today] || history['region'][yesterday];
        var lastUpdate = history['region'][today] ? today : yesterday;
        var sumCases = Object.values(dataProvince).reduce((a, b) => a + b, 0);
        document.getElementById('totalCases').innerHTML = sumCases.toLocaleString();
    }
});

currentDate = lastUpdate;
document.getElementById('dayFirst').onclick = function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var first = Object.keys(history[currentType])[0];
    lastUpdate = first;

    paintMap(currentType, history[currentType][first], first);
};
document.getElementById('dayBefore').onclick = function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var date = new Date(Date.parse(currentDate.innerHTML) - (24 * 60 * 60 * 1000));
    var before = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
    lastUpdate = before;

    paintMap(currentType, history[currentType][before], before);
};
document.getElementById('dayAfter').onclick = function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var date = new Date(Date.parse(currentDate.innerHTML) + (24 * 60 * 60 * 1000));
    var after = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
    lastUpdate = after;

    paintMap(currentType, history[currentType][after], after);
};
document.getElementById('dayLast').onclick = function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var last = Object.keys(history[currentType])[Object.keys(history[currentType]).length - 1];
    lastUpdate = last;

    paintMap(currentType, history[currentType][last], last);
};
