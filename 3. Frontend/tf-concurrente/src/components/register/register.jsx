import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./register.css";
import { withRouter } from "react-router-dom";

class Register extends Component {
	state = {
		isAuthenticated: "",
		userName: "",
		password: "",
		confirmPassword: "",
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
		localStorage.setItem("auth", true);
		localStorage.setItem("username", this.state.userName);
		this.setState(
			{
				isAuthenticated: true,
				userName: this.state.userName,
			},
			() => {
				this.props.onLogIn(this.state.isAuthenticated, this.state.userName);
				this.props.history.push("/predict");
			}
		);
	};

	render() {
		return (
			<React.Fragment>
				<div className="styleForForm">
					<h5>Registro</h5>
				</div>
				<Form className="styleForForm">
					<Form.Row className="row">
						<Col>
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
							<InputGroup className="mb-2">
								<Form.Control
									type="text"
									placeholder="Email"
									name="userName"
									value={this.state.userName}
									onChange={this.handleChange}
								/>
							</InputGroup>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<InputGroup className="mb-2">
								<Form.Control
									type="password"
									placeholder="Password"
									name="password"
									value={this.state.password}
									onChange={this.handleChange}
									className={
										this.state.password == this.state.confirmPassword
											? ""
											: "form-control is-invalid"
									}
								/>
							</InputGroup>
						</Col>
					</Form.Row>
					<Form.Row className="row">
						<Col>
							<InputGroup className="mb-2">
								<Form.Control
									type="password"
									placeholder="Confirm Password"
									name="confirmPassword"
									value={this.state.confirmPassword}
									onChange={this.handleChange}
									className={
										this.state.password == this.state.confirmPassword
											? ""
											: "form-control is-invalid"
									}
								/>
							</InputGroup>
						</Col>
					</Form.Row>
					{this.state.password != this.state.confirmPassword && (
						<div class="alert alert-danger" role="alert">
							The passwords don't match!
						</div>
					)}
				</Form>
				<div className="centerdiv">
					<Button
						variant="primary"
						className="button"
						onClick={this.submitForm}
					>
						Register
					</Button>{" "}
				</div>
			</React.Fragment>
		);
	}
}

export default withRouter(Register);
