// [START maps_marker_simple]

let map;
let markers = [];


// function initMap() {

//   console.log(data);

//   const myLatLng = { lat: data[0].Latitude, lng: data[0].Longitude };
//   map = new google.maps.Map(document.getElementById("map"), {
//     zoom: 15,
//     center: myLatLng,
//   });
//   map.addListener("click", (event) => {
//     addMarker(event.latLng, "Event Click");
//   });
//   addMarker(myLatLng);
//   const markers = data.forEach(node => {
//     let myLatLngLoop = { lat: node.Latitude, lng: node.Longitude };
//     addMarker(myLatLngLoop, node.Name)
//   });
// }

function initMap() {
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 3,
    center: { lat: 0, lng: -180 },
    mapTypeId: "terrain",
  });
  const flightPlanCoordinates = [];

  for (i = 0; i < data.length; i++) {
    flightPlanCoordinates.push({ lat: data[i].Latitude, lng: data[i].Longitude });
  }
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