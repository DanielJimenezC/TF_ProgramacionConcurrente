import React, { Component } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import { Row, Col, Grid, InputGroup } from "react-bootstrap";
import "./login.css";

class Login extends Component {
	state = {};
	render() {
		return (
			<React.Fragment>
				<Form className="styleForForm">
					<Form.Row className="row">
						<Col>
							<Form.Label>Username</Form.Label>
							<InputGroup className="mb-2">
								<Form.Control type="text" placeholder="Username" />
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
					<Button variant="primary" className="button">
						Login
					</Button>{" "}
				</div>
			</React.Fragment>
		);
	}
}

export default Login;
