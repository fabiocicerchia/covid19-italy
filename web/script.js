endpoints = {
    'region': '/limits_IT_regions.geojson',
    'province': '/limits_IT_provinces.geojson'
};

normalisedPlaces = {
    // italian
    'valle d\'aosta/vallée d\'aoste': 'valle d\'aosta',
    'bolzano/bozen':                  'bolzano',
    'massa-carrara':                  'massa carrara',
    'alto adige':                     'p.a. bolzano',
    'trentino':                       'p.a. trento',
    'emilia-romagna':                 'emilia romagna',
    // TODO: IT SHOULD INCLUDE ALTO ADIGE AS WELL (EITHER SUM OR SPLIT)
    'trentino-alto adige/südtirol':   'p.a. trento',
    'friuli-venezia giulia':          'friuli venezia giulia',
    'aosta':                          'valle d\'aosta',

    // english
    'florence':       'firenze',
    'genoa':          'genova',
    'mantua':         'mantova',
    'milan':          'milano',
    'naples':         'napoli',
    'padua':          'padova',
    'rome':           'roma',
    'south sardinia': 'sud sargegna',
    'syracuse':       'siracusa',
    'turin':          'torino',
    'venice':         'venezia',

    'latium':      'lazio',
    'apulia':      'puglia',
    'sardinia':    'sardegna',
    'sicily':      'sicilia',
    'south tyrol': 'p.a. trento'
};

tplPopup = {
    'province': '<strong>{PLACE}{REGION_LINK}</strong><br>Cases: {CASES} {TREND}<br><small>Updated on {DATE}</small>',
    'region': '<strong>{PLACE}</strong><br>Cases: {CASES} {TREND}<br>Recovered: {RECOVER}<br>Deaths: {DEATH}<br>Mortality Rate: {RATE}%<br><small>Updated on {DATE}</small>'
};

// UTILITIES
function heatMapColorforValue(value) {
    var h = (1.0 - value) * 240
    return "hsl(" + h + ", 100%, 50%)";
}

function xhrCall(path, callback) {
    var httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === 4 && httpRequest.status === 200) {
            callback(httpRequest.responseText);
        }
    };
    httpRequest.open('GET', path);
    httpRequest.send();
}

function fetchJSONFile(path, callback) {
    xhrCall(path, function (responseText) {
        callback(JSON.parse(responseText));
    });
}

function switchRegion(el) {
    document.getElementById('type-form-region').checked = true
    document.getElementById('value-form').value = normalisePlace(el.innerHTML);
    lookupAndRepaint();
}

function normalisePlace(place) {
    var lowerPlace = (typeof place !== 'undefined' ? place : '').toLowerCase();

    return typeof normalisedPlaces[lowerPlace] !== 'undefined' ? normalisedPlaces[lowerPlace] : lowerPlace;
}

geojsonCache = {region: undefined, province: undefined};

function fetchGeojson(type, callback) {
    if (geojsonCache[type] !== undefined) {
        callback(geojsonCache[type]);
        return;
    }

    fetchJSONFile(endpoints[type], function(items) {
        geojsonCache[type] = items;
        callback(items);
    });
}

function calcDateBefore(baseDate) {
    var date = new Date(Date.parse(baseDate) - (24 * 60 * 60 * 1000));
    var before = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');

    return before;
}

function calcDateAfter(baseDate) {
    var date = new Date(Date.parse(baseDate) + (24 * 60 * 60 * 1000));
    var after = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');

    return after;
}

function getDataPoint(date) {
    return typeof dataHistory[getCurrentType()][date] !== 'undefined' ? dataHistory[getCurrentType()][date] : undefined;
}

function getCurrentType() {
    return document.querySelector('input[name="type"]:checked').value;
}

function lookup() {
    var valueField = document.getElementById('value-form');
    valueField.value = valueField.value ? normalisePlace(valueField.value) : '';

    xhrCall('/trend.php?type=' + getCurrentType() + '&value=' + valueField.value, function (responseText) {
        document.getElementById('results').innerHTML = responseText;
    });
};
function lookupAndRepaint() {
    lookup();
    paintMap(getDataPoint(lastUpdate), getDataPoint(calcDateBefore(lastUpdate)), lastUpdate);
};

function paintMap(data, previousData, lastUpdateDate, type) {
    type = typeof type !== 'undefined' ? type : getCurrentType();

    document.getElementById('loader').classList.remove('d-none');
    document.getElementById('lastUpdate').innerHTML = lastUpdateDate;
    var max = Math.max.apply(Math, Object.values(data).map((i) => i.total));

    fetchGeojson(type, function (items) {
        if (geoLayer !== undefined) {
            geoLayer.remove();
        }

        geoLayer = L.geoJSON(items, {
           weight: 1,
           style: function(feature) {
                itemName = type === 'region'
                    ? normalisePlace(feature.properties.reg_name)
                    : normalisePlace(feature.properties.prov_name);

               var cases = data[itemName].total;
               if (cases > 0) {
                   return {
                       color: heatMapColorforValue(Math.min(cases / max + 0.75, 1))
                   };
               }
            },
            onEachFeature: function(feature, layer) {
                itemName = type === 'region'
                    ? normalisePlace(feature.properties.reg_name)
                    : normalisePlace(feature.properties.prov_name);

                var cases = data[itemName].total;
                var trendIcon = previousData == undefined
                    ? ''
                    : (cases > previousData[itemName].total ? ' &uarr;' : ' &darr;');

                var popupContent = tplPopup[type]
                    .replace('{PLACE}', itemName)
                    .replace('{REGION_LINK}', (type !== 'province' ? '' : ' (<a onclick="switchRegion(this)">'+feature.properties.reg_name+'</a>)'))
                    .replace('{CASES}', cases.toLocaleString())
                    .replace('{TREND}', trendIcon)
                    .replace('{RECOVER}', data[itemName].recovered.toLocaleString())
                    .replace('{DEATH}', data[itemName].death.toLocaleString())
                    .replace('{RATE}', (100 / cases * data[itemName].death).toFixed(0))
                    .replace('{DATE}', lastUpdateDate);

                if (itemName === 'p.a. trento') {
                    itemName = 'p.a. bolzano';
                    trendIcon = previousData == undefined
                        ? ''
                        : (cases > previousData[itemName].total ? ' &uarr;' : ' &darr;');
    
                    popupContent += '<br>'+tplPopup[type]
                        .replace('{PLACE}', 'P.A. Bolzano')
                        .replace('{REGION_LINK}', (type !== 'province' ? '' : ' (<a onclick="switchRegion(this)">'+feature.properties.reg_name+'</a>)'))
                        .replace('{CASES}', data[itemName].total.toLocaleString())
                        .replace('{TREND}', trendIcon)
                        .replace('{RECOVER}', data[itemName].recovered.toLocaleString())
                        .replace('{DEATH}', data[itemName].death.toLocaleString())
                        .replace('{RATE}', (100 / cases * data[itemName].death).toFixed(0))
                        .replace('{DATE}', lastUpdateDate);
                }
                layer.bindPopup(popupContent);
           }
        })
        .on('add', function () {
            document.getElementById('loader').classList.add('d-none');
        })
        .on('click', function(e) {
            document.getElementById('type-form-' + type).checked = true
            document.getElementById('value-form').value = type === 'region'
                    ? normalisePlace(e.sourceTarget.feature.properties.reg_name)
                    : normalisePlace(e.sourceTarget.feature.properties.prov_name);
            lookup();
        })
        .addTo(map);
    });
}

// SEARCH FORM
document.getElementById('search-form').onsubmit = function(e) {
    e.preventDefault();
    lookupAndRepaint();
};

// INIT DATA
geoLayer     = undefined;
lastUpdate   = undefined;
playInterval = undefined;
dataHistory  = {region: {}, province: {}};

// GEOMAP
var map = L.map('map', {
    zoomDelta: 0.25,
    zoomSnap: 0
}).setView([41.8719, 12.5674], 5.75);

L.tileLayer('https://api.mapbox.com/styles/v1/{id}/tiles/{z}/{x}/{y}?access_token=pk.eyJ1IjoiZmFiaW9jaWNlcmNoaWEiLCJhIjoiY2s3b2phNXhxMDlzbTNncDkzY3pkb2YxYSJ9.37ZdJEamTJMHvbticJ25CQ', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, ' +
        '<a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, ' +
        'Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    id: 'mapbox/light-v9'
}).addTo(map);

Papa.parse('/dpc-covid19-ita-province.csv', {
    header: true,
    download: true,
    complete: function(results) {
        var date = new Date();
        var today = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
        var yesterday = calcDateBefore(today);
        var beforeYesterday = calcDateBefore(yesterday);
        results.data.forEach(function(item) {
            if (item.data !== "") {
                var parsedDate = item.data.substr(0, 10);
                if (typeof dataHistory['province'][parsedDate] === 'undefined') dataHistory['province'][parsedDate] = {};
                province = normalisePlace(item.denominazione_provincia);
                if (typeof dataHistory['province'][parsedDate][province] !== "undefined") dataHistory['province'][parsedDate][province].total += parseInt(item.totale_casi, 10);
                else dataHistory['province'][parsedDate][province] = {
                    total:     parseInt(item.totale_casi, 10),
                    recovered: 0,
                    death:     0
                };
            }
        });
        var dataProvince  = dataHistory['province'][today] || dataHistory['province'][yesterday];
        var dataProvinceY = dataHistory['province'][today] ? dataHistory['province'][yesterday] : dataHistory['province'][beforeYesterday];
        lastUpdate        = dataHistory['province'][today] ? today : yesterday;
        var total         = Object.values(dataProvince).map((i) => i.total).reduce((a, b) => a + b, 0);
        var totalY        = Object.values(dataProvinceY).map((i) => i.total).reduce((a, b) => a + b, 0);
        var delta         = 100 / totalY * total;
        document.getElementById('num_total').innerHTML  = total.toLocaleString();
        if (delta > 100) document.getElementById('num_total_delta').innerHTML = '+' + (delta - 100).toFixed(0) + '%';
        else document.getElementById('num_total_delta').innerHTML = '-' + (delta * -1).toFixed(0) + '%';

        paintMap(dataProvince, getDataPoint('province', calcDateBefore(lastUpdate)), lastUpdate, 'province');
    }
});

Papa.parse('/dpc-covid19-ita-regioni.csv', {
    header: true,
    download: true,
    complete: function(results) {
        var date = new Date();
        var today = date.getFullYear() + '-' + (date.getMonth() + 1).toString().padStart(2, '0') + '-' + (date.getDate()).toString().padStart(2, '0');
        var yesterday = calcDateBefore(today);
        var beforeYesterday = calcDateBefore(yesterday);

        results.data.forEach(function(item) {
            var parsedDate = item.data.substr(0, 10);
            if (typeof dataHistory['region'][parsedDate] === 'undefined') dataHistory['region'][parsedDate] = {};
            region = normalisePlace(item.denominazione_regione);
            dataHistory['region'][parsedDate][region] = {
                total:        parseInt(item.totale_casi, 10),
                hospitalised: parseInt(item.ricoverati_con_sintomi, 10),
                icu:          parseInt(item.terapia_intensiva, 10),
                home:         parseInt(item.isolamento_domiciliare, 10),
                total_active: parseInt(item.totale_positivi, 10),
                //new_active:   parseInt(item.nuovi_positivi, 10),
                new_active:   parseInt(item.variazione_totale_positivi, 10),
                recovered:    parseInt(item.dimessi_guariti, 10),
                death:        parseInt(item.deceduti, 10),
                tests:        parseInt(item.tamponi, 10),
            };
        });
        var dataRegion  = dataHistory['region'][today] || dataHistory['region'][yesterday];
        var dataRegionY = dataHistory['region'][today] ? dataHistory['region'][yesterday] : dataHistory['region'][beforeYesterday];
        var recovered   = Object.values(dataRegion).map((i) => i.recovered).reduce((a, b) => a + b, 0);
        var active      = Object.values(dataRegion).map((i) => i.total_active).reduce((a, b) => a + b, 0);
        var activeY     = Object.values(dataRegionY).map((i) => i.total_active).reduce((a, b) => a + b, 0);
        var newCases    = Object.values(dataRegion).map((i) => i.new_active).reduce((a, b) => a + b, 0);
        var newCasesY   = Object.values(dataRegionY).map((i) => i.new_active).reduce((a, b) => a + b, 0);
        var deaths      = Object.values(dataRegion).map((i) => i.death).reduce((a, b) => a + b, 0);
        var sumCases    = Object.values(dataRegion).map((i) => i.total).reduce((a, b) => a + b, 0);
        var activeDelta = 100 / activeY * active;
        var newDelta    = 100 - (100 / newCasesY * newCases);
        document.getElementById('num_new').innerHTML          = newCases.toLocaleString();
        document.getElementById('num_active').innerHTML       = active.toLocaleString();
        document.getElementById('num_death').innerHTML        = deaths.toLocaleString();
        document.getElementById('num_death_rate').innerHTML   = (100 / sumCases * deaths).toFixed(0) + '%';
        document.getElementById('num_recover').innerHTML      = recovered.toLocaleString();
        document.getElementById('num_recover_rate').innerHTML = (100 / sumCases * recovered).toFixed(0) + '%';
        if (activeDelta > 100) document.getElementById('num_active_delta').innerHTML = (activeDelta - 100).toFixed(0) + '%';
        else document.getElementById('num_active_delta').innerHTML = '-' + (activeDelta * -1).toFixed(0) + '%';
        if (newDelta > 100) document.getElementById('num_new_delta').innerHTML = '-' + (newDelta - 100).toFixed(0) + '%';
        else document.getElementById('num_new_delta').innerHTML = (newDelta * -1).toFixed(0) + '%';

        var dataLabels = Object.keys(dataHistory['region']).filter((i) => i !== "");
        var newPoints = Object.values(dataHistory['region']).filter((i) => Object.values(i).length > 1).map((v) => Object.values(v).map((i) => i.new_active || 0).reduce((a, b) => a + b, 0));
        var ctx = document.getElementById('chart').getContext('2d');
        var myLineChart = new Chart(ctx, {
            type: 'line',
            options: {
            },
            data: {
                labels: dataLabels,
                datasets: [
                    {label: 'New Cases', data: newPoints, borderColor: '#de7119', backgroundColor: 'transparent'},
                ]
            }
        });
    }
});

document.getElementById('dayPlay').addEventListener('click', function() {
    document.getElementById('dayFirst').dispatchEvent(new MouseEvent('click', {}));
    playInterval = setInterval(function() {
        document.getElementById('dayAfter').dispatchEvent(new MouseEvent('click', {}));
    }, 700);
});
document.getElementById('dayFirst').addEventListener('click', function() {
    lastUpdate = Object.keys(dataHistory[getCurrentType()])[0];

    paintMap(getDataPoint(lastUpdate), undefined, lastUpdate);
});
document.getElementById('dayBefore').addEventListener('click', function() {
    var before = calcDateBefore(lastUpdate);
    if ((dataPoint = getDataPoint(before)) !== undefined) {
        lastUpdate = before;

        paintMap(dataPoint, getDataPoint(calcDateBefore(before)), before);
    }
});
document.getElementById('dayAfter').addEventListener('click', function() {
    var after = calcDateAfter(lastUpdate);

    if ((dataPoint = getDataPoint(after)) !== undefined) {
        lastUpdate = after;

        paintMap(dataPoint, getDataPoint(calcDateBefore(after)), after);
    } else {
        // stop the play, even if didn't start
        playInterval = undefined;
    }
});
document.getElementById('dayLast').addEventListener('click', function() {
    lastUpdate = Object.keys(dataHistory[getCurrentType()])[Object.keys(dataHistory[currentType]).length - 1];

    paintMap(getDataPoint(lastUpdate), getDataPoint(calcDateBefore(lastUpdate)), lastUpdate);
});
document.querySelectorAll('.btn-switch[data-type]').forEach(function (item) {
    item.addEventListener('click', function() {
        document.querySelector('#type-province input').checked = false;
        document.querySelector('#type-region input').checked = false;
        document.querySelector('#type-' + this.dataset['type'] + ' input').checked = true;
        paintMap(getDataPoint(lastUpdate), getDataPoint(calcDateBefore(lastUpdate)), lastUpdate);
    });
});

$(function () {
    $('[data-toggle="tooltip"]').tooltip()
});
