import './Header.scss';
import React, {Component} from "react";

class Header extends Component {
    render() {
        return (
            <div className="header">
                <div className="column">
                    <h2>Chat-App</h2>
                </div>
                <div className="input-field column">
                    <input placeholder="Enter key" id="search_string" type="text" className="validate"
                           onChange={e => this.props.keyChange(e.target.value)} value={this.props.key}
                    />
                </div>
            </div>
        );
    }
}

export default Header;
