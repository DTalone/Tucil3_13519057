// [START maps_marker_simple]

let map;
let markers = [];
let idxClick = 0;
let path = [];
let distance = 0;

function initMap() {
  const flightPlanCoordinates = [];

  for (i = 0; i < data.length; i++) {
    flightPlanCoordinates.push({ lat: data[i].Latitude, lng: data[i].Longitude });
    addMarker(flightPlanCoordinates[i], data[i].Name)
  }
  map = new google.maps.Map(document.getElementById("map"), {
    zoom: 15,
    center: flightPlanCoordinates[0],
  });
  map.addListener("click", (event) => {
    addMarker(event.latLng, "Event Click " + idxClick);
    idxClick++;
  });
  setMapOnAll(map);
}

function initRute() {
  const flightPlanCoordinates = [];
  distance = data[data.length - 1].Latitude;
  if (distance == 0) {
    console.log("YAMAMAMAM")
    var arrayy = ["1", "2", "3"];
    var dist = document.getElementById("distance")
    for (j = 0; j < arrayy.length; j++) {
      dist.innerHTML += `<li>${arrayy[j]}</li>`
    }
    // document.getElementById("distance").innerHTML = `Tidak Ada apa Jalur</a>`;
    initMap();
    return
  } else {
    document.getElementById("distance").innerHTML = "Distancee : " + distance;
  }
  for (i = 0; i < data.length - 1; i++) {
    flightPlanCoordinates.push({ lat: data[i].Latitude, lng: data[i].Longitude });
    addMarker(flightPlanCoordinates[i], data[i].Name)
  }
  map = new google.maps.Map(document.getElementById("map"), {
    zoom: 15,
    center: flightPlanCoordinates[0],
  });
  map.addListener("click", (event) => {
    addMarker(event.latLng, "Event Click " + idxClick);
    idxClick++;
  });
  const flightPath = new google.maps.Polyline({
    path: flightPlanCoordinates,
    geodesic: true,
    strokeColor: "#FF0000",
    strokeOpacity: 1.0,
    strokeWeight: 2,
  });
  flightPath.setMap(map);
}

// Adds a marker to the map and push to the array.
function addMarker(location, label) {
  const marker = new google.maps.Marker({
    position: location,
    map: map,
  });
  const infowindow = new google.maps.InfoWindow({
    content: "<j>" + label + "</h>" + "<p>Location:" + marker.getPosition() + "</p>",
  });
  marker.addListener("click", () => {
    infowindow.open(map, marker);
  });
  markers.push(marker);
}

// Sets the map on all markers in the array.
function setMapOnAll(map) {
  for (let i = 0; i < markers.length; i++) {
    markers[i].setMap(map);
  }
}

// Removes the markers from the map, but keeps them in the array.
function clearMarkers() {
  setMapOnAll(null);
}

// Shows any markers currently in the array.
function showMarkers() {
  setMapOnAll(map);
}

// Deletes all markers in the array by removing references to them.
function deleteMarkers() {
  clearMarkers();
  markers = [];
}