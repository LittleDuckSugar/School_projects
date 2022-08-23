import React, { useState } from 'react';
import Input from '../../components/Input';
import userService from "../../services/user.service";
import Button from "../../components/Button";
import { useRouter } from "next/router";

const Index = () => {

	const [user, setUser] = useState();

	const router = useRouter();

	const submitLogin = (e) => {
		e.preventDefault();
		userService.login(user)
			.then((data) => {
				console.log(data);
				localStorage.setItem('token', data.jwt);
				router.push('/profil')
			})
			.catch(err => console.log(err))
	}

	return (
		<div className="page__login">
			<form className="form" onSubmit={(e) => submitLogin(e)}>
				<Input
					type="email"
					label="Email"
					placeholder="Veuillez saisir votre adresse email"
					name="email"
					id="email"
					required={true}
					classes="form__input"
					handleChange={(e) => setUser({ ...user, identifier: e.target.value })}
				/>
				<Input
					type="password"
					label="Password"
					placeholder="Veuillez saisir votre mot de passe"
					name="password"
					id="password"
					required={true}
					classes="form__input"
					handleChange={(e) => setUser({ ...user, password: e.target.value })}
				/>
				<Button title="envoyer" classes="btn btn__color-black" type="submit" />
			</form>
		</div>
	);
}

export default Index;
