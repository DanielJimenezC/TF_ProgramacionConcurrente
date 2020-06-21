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
		this.state = { userName: "" };
	}

	handleLogIn(newName) {
		this.setState({
			userName: newName,
		});
	}

	render() {
		return (
			<React.Fragment>
				<Router>
					<Layout userName={this.state.userName} />
					<Switch>
						<Route
							exact
							path="/login"
							render={(props) => <Login onLogIn={this.handleLogIn} />}
						/>
						<PrivateRoute path="/predict" component={Predict} />
					</Switch>
				</Router>
			</React.Fragment>
		);
	}
}

export default App;
