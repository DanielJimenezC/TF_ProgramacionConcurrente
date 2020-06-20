import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./login.css";
import {
	BrowserRouter as Router,
	Route,
	Switch,
	Redirect,
} from "react-router-dom";
import { toJS } from "mobx";

class Login extends Component {
	state = {
		isAuthenticated: false,
		userName: "",
	};

	handleChange = (event) => {
		const target = event.target;
		const value = target.value;
		const name = target.name;

		this.setState({
			[name]: value,
		});
	};
	submitForm = () => {
		this.state.isAuthenticated = true;
		this.setState({
			isAuthenticated: true,
		});
		this.props.history.push("/predict", {
			responseState: this.state,
		});
	};

	render() {
		return (
			<React.Fragment>
				<Form className="styleForForm">
					<Form.Row className="row">
						<Col>
							<Form.Label>Username</Form.Label>
							<InputGroup className="mb-2">
								<Form.Control
									type="text"
									placeholder="Username"
									name="userName"
									value={this.state.userName}
									onChange={this.handleChange}
								/>
							</InputGroup>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<Form.Label>Password</Form.Label>
							<InputGroup className="mb-2">
								<Form.Control type="password" placeholder="Password" />
							</InputGroup>
						</Col>
					</Form.Row>
				</Form>
				<div className="centerdiv">
					<Button
						variant="primary"
						className="button"
						onClick={this.submitForm}
					>
						Login
					</Button>{" "}
				</div>
			</React.Fragment>
		);
	}
}

export default Login;
