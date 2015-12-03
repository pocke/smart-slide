// TODO: localhost
var conn = new WebSocket("ws://localhost:" + port.toString());
conn.onmessage = function () {
  slideshow.gotoNextSlide();
};
