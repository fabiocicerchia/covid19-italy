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

function switchRegion(el) {
    document.getElementById('type-form-region').checked = true
    document.getElementById('value-form').value = normalisePlace(el.innerHTML);
    document.getElementById('search-form').submit();
}

function normalisePlace(place) {
    var lowerPlace = place === undefined ? '' : place.toLowerCase();

    // italian
    if (lowerPlace === 'valle d\'aosta/vallée d\'aoste') return 'valle d\'aosta';
    else if (lowerPlace === 'bolzano/bozen') return 'bolzano';
    else if (lowerPlace === 'massa-carrara') return 'massa carrara';
    else if (lowerPlace === 'alto adige') return 'p.a. bolzano';
    else if (lowerPlace === 'trentino') return 'p.a. trento';
    else if (lowerPlace === 'emilia-romagna') return 'emilia romagna';
    else if (lowerPlace === 'trentino-alto adige/südtirol') return 'p.a. trento'; // TODO: IT SHOULD INCLUDE ALTO ADIGE AS WELL (EITHER SUM OR SPLIT)
    else if (lowerPlace === 'friuli-venezia giulia') return 'friuli venezia giulia';
    else if (lowerPlace === 'aosta') return 'valle d\'aosta';

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

function paintMap(type, data, lastUpdateDate) {
    document.getElementById('lastUpdate').innerHTML = lastUpdateDate;
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
               var popupContent = '<strong>' + place + (type !== 'province' ? '' : ' (<a onclick="switchRegion(this)">'+feature.properties.reg_name+'</a>)') + '</strong><br>';
               popupContent += 'Cases: ' + cases + '<br>';
               popupContent += '<small>Updated on ' + lastUpdateDate + '</small>';

	       if (place === 'p.a. trento') {
                   popupContent += '<br><strong>P.A. Bolzano</strong><br>';
                   popupContent += 'Cases: ' + data['p.a. bolzano'] + '<br>';
                   popupContent += '<small>Updated on ' + lastUpdateDate + '</small>';
	       }
               layer.bindPopup(popupContent);
           }
        })
        .addTo(map)
        .on('click', function(e) {
            document.getElementById('type-form-' + type).checked = true
            document.getElementById('value-form').value = type === 'region'
                    ? normalisePlace(e.sourceTarget.feature.properties.reg_name)
                    : normalisePlace(e.sourceTarget.feature.properties.prov_name);
            document.getElementById('search-form').submit();
        });
    });
}

// SEARCH FORM
document.getElementById('search-form').onsubmit = function(e) {
    var valueField = document.getElementById('value-form');
    valueField.value = valueField.value ? normalisePlace(valueField.value) : '';
    var currentType = document.querySelector('input[name="type"]:checked').value;
    paintMap(currentType, dataHistory[currentType][lastUpdate], lastUpdate);
};

// COVID DATA
dataHistory = {region: {}, province: {}};

// GEOMAP
var map = L.map('map').setView([41.8719, 12.5674], 6);

L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token=pk.eyJ1IjoiZmFiaW9jaWNlcmNoaWEiLCJhIjoiY2s3b2phNXhxMDlzbTNncDkzY3pkb2YxYSJ9.37ZdJEamTJMHvbticJ25CQ', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, ' +
        '<a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
        'Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    id: 'mapbox/light-v9'
}).addTo(map);

geoLayer = undefined;
lastUpdate = undefined;

Papa.parse('/dpc-covid19-ita-province.csv', {
    header: true,
    download: true,
    complete: function(results) {
        var date = new Date();
        var today = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
        var yesterday = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate() - 1).toString().padStart(2, '0');
        results.data.forEach(function(item) {
            if (item.data !== "") {
                var parsedDate = item.data.substr(0, 10);
                if (typeof dataHistory['province'][parsedDate] === 'undefined') dataHistory['province'][parsedDate] = {};
                province = normalisePlace(item.denominazione_provincia);
                dataHistory['province'][parsedDate][province] = parseInt(item.totale_casi, 10);
            }
        });
        var dataProvince = dataHistory['province'][today] || dataHistory['province'][yesterday];
        lastUpdate = dataHistory['province'][today] ? today : yesterday;
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
            if (typeof dataHistory['region'][parsedDate] === 'undefined') dataHistory['region'][parsedDate] = {};
	    region = normalisePlace(item.denominazione_regione);
            dataHistory['region'][parsedDate][region] = parseInt(item.totale_casi, 10);
        });
        var dataRegion = dataHistory['region'][today] || dataHistory['region'][yesterday];
        var sumCases = Object.values(dataRegion).reduce((a, b) => a + b, 0);
        document.getElementById('totalCases').innerHTML = sumCases.toLocaleString();
    }
});

currentDate = lastUpdate;
document.getElementById('dayFirst').addEventListener('click', function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var first = Object.keys(dataHistory[currentType])[0];
    lastUpdate = first;

    paintMap(currentType, dataHistory[currentType][first], first);
});
document.getElementById('dayBefore').addEventListener('click', function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var date = new Date(Date.parse(currentDate.innerHTML) - (24 * 60 * 60 * 1000));
    var before = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');

    if (dataHistory[currentType][before] !== undefined) {
        lastUpdate = before;

        paintMap(currentType, dataHistory[currentType][before], before);
    }
});
document.getElementById('dayAfter').addEventListener('click', function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var date = new Date(Date.parse(currentDate.innerHTML) + (24 * 60 * 60 * 1000));
    var after = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');

    if (dataHistory[currentType][after] !== undefined) {
        lastUpdate = after;

        paintMap(currentType, dataHistory[currentType][after], after);
    }
});
document.getElementById('dayLast').addEventListener('click', function() {
    var currentType = document.querySelector('input[name="type"]:checked').value;
    var last = Object.keys(dataHistory[currentType])[Object.keys(dataHistory[currentType]).length - 1];
    lastUpdate = last;

    paintMap(currentType, dataHistory[currentType][last], last);
});
document.getElementById('type-region').addEventListener('click', function() {
    paintMap('region', dataHistory['region'][lastUpdate], lastUpdate);
});
document.getElementById('type-province').addEventListener('click', function() {
    paintMap('province', dataHistory['province'][lastUpdate], lastUpdate);
});
