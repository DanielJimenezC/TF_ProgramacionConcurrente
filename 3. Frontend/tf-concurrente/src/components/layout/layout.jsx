import React, { Component } from "react";
import { Navbar, Nav } from "react-bootstrap";

class Layout extends Component {
	state = {
		isAuthenticated: false,
		userName: "",
	};

	render() {
		return (
			<React.Fragment>
				<Navbar bg="light" expand="lg" bg="dark" variant="dark">
					<Navbar.Brand>Smart Rona</Navbar.Brand>
					<Navbar.Toggle aria-controls="basic-navbar-nav" />
					<Navbar.Collapse className="justify-content-end">
						<Nav className="mr-auto">
							{this.props.state.isAuthenticated && (
								<Nav.Link href="predict">Examen de Covid-19</Nav.Link>
							)}
						</Nav>
						<Navbar.Text>
							Signed in as: <a href="#login">{this.props.state.userName}</a>
						</Navbar.Text>
					</Navbar.Collapse>
				</Navbar>
			</React.Fragment>
		);
	}
}

export default Layout;
