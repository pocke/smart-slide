// TODO: localhost
var conn = new WebSocket("ws://" + location.hostname + ":" + port.toString());
conn.onmessage = function () {
  slideshow.gotoNextSlide();
};
