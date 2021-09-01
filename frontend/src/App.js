import './App.css';
import React, {Component} from "react";
import {connect, sendMsg} from "./api";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory";
import ChatInput from "./components/ChatInput";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      key: "React",
      chatHistory: []
    };
  }

  componentDidMount() {
    connect((msg) => {
      this.setState(prevState => ({
        chatHistory: [...prevState.chatHistory, msg]
      }));
    });
  }

  send(event) {
    if(event.keyCode === 13){
      sendMsg(1, event.target.value);
      event.target.value = "";
    }
  }

  changeKey(key) {
    // this.setState(({
    //   key: key
    // }));
    sendMsg(0, key)
  }

  render() {
    return (
        <div className="App">
          <Header keyChange={this.changeKey} key={this.state.key}/>
          <ChatHistory chatHistory={this.state.chatHistory} />
          <ChatInput send={this.send}/>
        </div>
    );
  }
}

export default App;
