import React, { useState, useEffect } from 'react';
import TitlePage from '../../components/TitlePage';
import userService from '../../services/user.service';
import withAuth from '../../HOC/withAuth';
import Button from "../../components/Button";
const Index = () => {

	const [user, setUser] = useState();

	useEffect(() => {
		userService.getMe(localStorage.getItem('token'))
			.then(data => {
				console.log(data);
				setUser(data);
			})
			.catch(err => console.log(err))
	}, [])

	return (
		<div>
			<TitlePage title="Profil" />
			<div className="text__center">
				<p>Nom: {user && user.firstName}</p>
				<p>Prénom: {user && user.lastName}</p>
				<p>Email: {user && user.email}</p>
				<p>
					<Button title="logout" classes="btn btn__color-white" />
				</p>
			</div>
		</div>
	);
}

export default withAuth(Index);
