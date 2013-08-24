function initialize_locations() {
  var mapOptions = {
    center: new google.maps.LatLng(32.7833, -79.9333),
    zoom: 13,
    mapTypeId: google.maps.MapTypeId.ROADMAP
  };

  var input = /** @type {HTMLInputElement} */(document.getElementById('create_outing_location'));
  var autocomplete = new google.maps.places.Autocomplete(input);

  var infowindow = new google.maps.InfoWindow();

  google.maps.event.addListener(autocomplete, 'place_changed', function() {
    infowindow.close();
    //input.className = '';
    var place = autocomplete.getPlace();
    if (!place.geometry) {
      // Inform the user that the place was not found and return.
      //input.className = 'notfound';
      return;
    }

    var address = '';
    if (place.address_components) {
      address = [
        (place.address_components[0] && place.address_components[0].short_name || ''),
        (place.address_components[1] && place.address_components[1].short_name || ''),
        (place.address_components[2] && place.address_components[2].short_name || '')
      ].join(' ');
    }

    infowindow.setContent('<div><strong>' + place.name + '</strong><br>' + address);
    alert(place.name);
  });
}

//google.maps.event.addDomListener(window, 'load', initialize_locations);