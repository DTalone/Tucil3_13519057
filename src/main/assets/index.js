// [START maps_marker_simple]



function initMap() {

  console.log("anjay");
  console.log(data);

  const myLatLng = { lat: data[0].Latitude, lng: data[0].Longitude };
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 4,
    center: myLatLng,
  });

  data.forEach(node => {
    let myLatLngLoop = { lat: node.Latitude, lng: node.Longitude };
    console.log(myLatLngLoop);
    new google.maps.Marker({
      position: myLatLngLoop,
      map,
      title: "Hello World!",
    });
  });
}
  // [END maps_marker_simple]