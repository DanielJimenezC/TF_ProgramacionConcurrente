import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import Modal from "react-bootstrap/Modal";
import "./login.css";
import { withRouter } from "react-router-dom";
import axios from "axios";

class Login extends Component {
	state = {
		isAuthenticated: "",
		username: "",
		password: "",
		valueOfModal: false,
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
		/*
		localStorage.setItem("auth", true);
		localStorage.setItem("username", this.state.username);
		this.setState(
			{
				isAuthenticated: true,
				username: this.state.username,
			},
			() => {
				this.props.onLogIn(this.state.isAuthenticated, this.state.username);
				this.props.history.push("/predict");
			}
		);
		*/
		axios
			.post("http://localhost:5000/api/users/login", this.state)
			.then((resonse) => {
				console.log(resonse.data);
				if (resonse.data.message == "Succesfull Login") {
					console.log("EXITOO");
					localStorage.setItem("auth", true);
					localStorage.setItem("username", this.state.username);
					this.setState(
						{
							isAuthenticated: true,
							username: this.state.username,
						},
						() => {
							this.props.onLogIn(
								this.state.isAuthenticated,
								this.state.username
							);
							this.props.history.push("/dashboard");
						}
					);
				} else {
					this.setState({
						valueOfModal: true,
					});
				}
			})
			.catch((error) => {
				console.log(error);
			});
	};

	render() {
		return (
			<React.Fragment>
				<div className="styleForForm">
					<h5
						style={{
							paddingTop: "20px",
						}}
					>
						Login
					</h5>
				</div>
				<Form className="styleForForm">
					<Form.Row className="row">
						<Col>
							<InputGroup className="mb-2">
								<Form.Control
									type="text"
									placeholder="Username"
									name="username"
									value={this.state.username}
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
								/>
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
				<Modal
					size="lg"
					aria-labelledby="contained-modal-title-vcenter"
					centered
					show={this.state.valueOfModal}
					onHide={() => this.setState({ valueOfModal: false })}
				>
					<Modal.Header closeButton>
						<Modal.Title id="contained-modal-title-vcenter">
							Modal heading
						</Modal.Title>
					</Modal.Header>
					<Modal.Body>
						<h4>Centered Modal</h4>
						<p>
							Cras mattis consectetur purus sit amet fermentum. Cras justo odio,
							dapibus ac facilisis in, egestas eget quam. Morbi leo risus, porta
							ac consectetur ac, vestibulum at eros.
						</p>
					</Modal.Body>
				</Modal>
			</React.Fragment>
		);
	}
}

export default withRouter(Login);
