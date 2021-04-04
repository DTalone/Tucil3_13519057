// [START maps_marker_simple]



function initMap() {

  console.log("anjay");
  console.log(data);

  const myLatLng = { lat: data[0].Latitude, lng: data[0].Longitude };
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 15,
    center: myLatLng,
  });
  map.addListener("click", (e) => {
    placeMarkerAndPanTo(e.latLng, map);
  });

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

function placeMarkerAndPanTo(latLng, map) {
  new google.maps.Marker({
    position: latLng,
    map: map,
  });
  map.panTo(latLng);
}