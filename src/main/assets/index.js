// [START maps_marker_simple]


let map;
let markers = [];

function initMap() {

  console.log("anjay");
  console.log(data);

  const myLatLng = { lat: data[0].Latitude, lng: data[0].Longitude };
  map = new google.maps.Map(document.getElementById("map"), {
    zoom: 15,
    center: myLatLng,
  });
  map.addListener("click", (event) => {
    addMarker(event.latLng);
  });
  addMarker(myLatLng);
  const markers = data.forEach(node => {
    let myLatLngLoop = { lat: node.Latitude, lng: node.Longitude };
    console.log(myLatLngLoop);
    const marker = new google.maps.Marker({
      position: myLatLngLoop,
      map,
    });
    const infowindow = new google.maps.InfoWindow({
      content: "<j>"+node.Name+"</h>"+"<p>Location:" + marker.getPosition() + "</p>",
    });
    marker.addListener("click", () => {
      infowindow.open(map, marker);
    });
  });
  // Add a marker clusterer to manage the markers.
  new MarkerClusterer(map, markers, {
    imagePath:
      "https://developers.google.com/maps/documentation/javascript/examples/markerclusterer/m",
  });
  
}

// Adds a marker to the map and push to the array.
function addMarker(location) {
  const marker = new google.maps.Marker({
    position: location,
    map: map,
  });
  const infowindow = new google.maps.InfoWindow({
    content: "<j>"+"Event click"+"</h>"+"<p>Location:" + marker.getPosition() + "</p>",
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