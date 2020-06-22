import React, { Component } from "react";
import Predict from "./predict/predict";
import {
	BrowserRouter as Router,
	Route,
	Switch,
	Redirect,
} from "react-router-dom";
import Layout from "./layout/layout";
import Login from "./login/login";

const PrivateRoute = ({ component: Component, ...rest }) => (
	<Route
		{...rest}
		render={(props) =>
			props.location.state.responseState.isAuthenticated === true ? (
				<Component {...props} />
			) : (
				<Redirect to="/login" />
			)
		}
	/>
);

class App extends Component {
	constructor(props) {
		super(props);
		this.handleLogIn = this.handleLogIn.bind(this);
		this.state = { userName: "", isAuthenticated: false };
	}

	handleLogIn(newState, newName) {
		this.setState({
			isAuthenticated: newState,
			userName: newName,
		});
	}

	render() {
		return (
			<React.Fragment>
				<Router>
					<Layout state={this.state} />
					<Switch>
						<Route
							exact
							path="/login"
							render={(props) => <Login onLogIn={this.handleLogIn} />}
						/>
						<Route
							exact
							path="/predict"
							render={(props) =>
								this.state.isAuthenticated === true ? (
									<Predict state={this.state} />
								) : (
									<Redirect to="/login" />
								)
							}
						/>
					</Switch>
				</Router>
			</React.Fragment>
		);
	}
}

export default App;
