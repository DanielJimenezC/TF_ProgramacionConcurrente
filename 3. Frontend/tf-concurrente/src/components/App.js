import React from "react";
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

function App() {
	return (
		<React.Fragment>
			<Router>
				<Layout />
				<Switch>
					<Route exact path="/login" component={Login} />
					<PrivateRoute path="/predict" component={Predict} />
				</Switch>
			</Router>
		</React.Fragment>
	);
}

export default App;
