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
import Register from "./register/register";
import Faq from "./faq/faq";
import Dashboard from "./dashboard/dashboard";
import RegisterData from "./registerdata/registerdata";
import Groups from "./groups/groups";

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
				<Layout state={this.state} />
				<Router>
					<Switch>
						<Route
							exact
							path="/"
							render={(props) => (
								<div>
									<Login onLogIn={this.handleLogIn} />
									<Register onLogIn={this.handleLogIn} />
								</div>
							)}
						/>
						<Route
							exact
							path="/predict"
							render={(props) =>
								Boolean(localStorage.getItem("auth")) === true ? (
									<Predict state={this.state} />
								) : (
									<Redirect to="/" />
								)
							}
						/>
						<Route
							exact
							path="/faq"
							render={(props) =>
								Boolean(localStorage.getItem("auth")) === true ? (
									<Faq state={this.state} />
								) : (
									<Redirect to="/" />
								)
							}
						/>
						<Route
							exact
							path="/dashboard"
							render={(props) =>
								Boolean(localStorage.getItem("auth")) === true ? (
									<Dashboard state={this.state} />
								) : (
									<Redirect to="/" />
								)
							}
						/>
						<Route
							exact
							path="/groups"
							render={(props) =>
								Boolean(localStorage.getItem("auth")) === true ? (
									<RegisterData state={this.state} />
								) : (
									<Redirect to="/" />
								)
							}
						/>
						<Route
							exact
							path="/register"
							render={(props) =>
								Boolean(localStorage.getItem("auth")) === true ? (
									<Groups state={this.state} />
								) : (
									<Redirect to="/" />
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
