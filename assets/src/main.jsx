const React = require('react');
const ReactDOM = require("react-dom");
const req = require('superagent');

const KeyButton = React.createClass({
  propTypes: {
    keyCode: React.PropTypes.string.isRequired,
  },
  post: function () {
    req
    .post("/key")
    .send(this.props.keyCode).end();
  },
  render: function () {
    return (
      <button onClick={this.post}>{this.props.keyCode}</button>
    );
  },
});

const App = React.createClass({
  render: function () {
    return (
      <div>
        <KeyButton keyCode='left'></KeyButton>
        <KeyButton keyCode='right'></KeyButton>
      </div>
    )
  },
})

ReactDOM.render(
  <App />,
  document.querySelector("#react-main")
)
