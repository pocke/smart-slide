// TODO: localhost
var conn = new WebSocket("ws://" + location.hostname + ":" + port.toString());
conn.onmessage = function (ev) {
  var text = JSON.parse(ev.data);
  if (text === 'right') {
    slideshow.gotoNextSlide();
  } else {
    slideshow.gotoPreviousSlide();
  }
};
