const React = require('react');
const ReactDOM = require("react-dom");
const req = require('superagent');

const KeyButton = React.createClass({
  propTypes: {
    keyCode: React.PropTypes.string.isRequired,
    color:   React.PropTypes.string.isRequired,
  },
  post: function () {
    req
    .post("/key")
    .send(this.props.keyCode).end((err, res) => {
      if (!res) {
        alert("Server is dead");
      }
      if (!res.ok) {
        alert("error!", res.text);
      }
    });
  },
  render: function () {
    const style = {
      display: "table-cell",
      width: "50%",
      height: "100%",
      'background-color': this.props.color,
      'text-align': 'center',
      'vertical-align': 'middle',
      cursor: 'pointer',
    };

    return (
      <div
        onClick={this.post}
        style={style}
      >{this.props.keyCode}</div>
    );
  },
});

const App = React.createClass({
  render: function () {
    return (
      <div style={{display: "table", width: "100%", height: "100%"}}>
        <KeyButton keyCode='left'  color='#9fc'></KeyButton>
        <KeyButton keyCode='right' color="#f9c"></KeyButton>
      </div>
    )
  },
})

ReactDOM.render(
  <App />,
  document.querySelector("#react-main")
)
