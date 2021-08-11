import './App.css';
import React, {Component} from "react";
import {connect, sendMsg} from "./api";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("hello");
    sendMsg("hello");
  }

  render() {
    return (
        <div className="App">
          <button onClick={this.send}>Hit</button>
        </div>
    );
  }
}

export default App;
