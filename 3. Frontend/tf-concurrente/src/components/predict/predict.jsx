import React, { Component } from "react";

class Predict extends Component {
	render() {
		return (
			<React.Fragment>
				<h1>HELLO {this.props.location.state.responseState.userName}</h1>
			</React.Fragment>
		);
	}
}

export default Predict;
