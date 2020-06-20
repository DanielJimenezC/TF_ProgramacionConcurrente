import React, { Component } from "react";
import { Navbar, Nav } from "react-bootstrap";

class Layout extends Component {
	render() {
		return (
			<React.Fragment>
				<Navbar bg="light" expand="lg" bg="dark" variant="dark">
					<Navbar.Brand>Smart Rona</Navbar.Brand>
					<Navbar.Toggle aria-controls="basic-navbar-nav" />
					<Navbar.Collapse className="justify-content-end">
						<Navbar.Text>
							Signed in as: <a href="#login">Mark Otto</a>
						</Navbar.Text>
					</Navbar.Collapse>
				</Navbar>
			</React.Fragment>
		);
	}
}

export default Layout;
