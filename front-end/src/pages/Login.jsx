// eslint-disable-next-line no-unused-vars
import React, {useEffect, useState} from 'react';
import './Login.css';
import EyeOpened from './assets/password.svg';
import EyeClosed from './assets/hidden-password.svg';
import { useNavigate } from 'react-router-dom';

function fetchUserData(email, password) {
	const payload = {
		action: "auth_user",
		auth: {
			email: email,
			password: password,
		}
	}

	const headers = new Headers();
	headers.append("Content-Type", "application/json");

	return fetch("http:\/\/localhost:8080/handle", {
		method: 'POST',
		body: JSON.stringify(payload),
		headers: headers,
	})
		.then(response => response.json())
		.then(data => {
			return data;
		})
		.catch(error => console.error(error));
}

function Login(props) {
	const [email, setEmail] = useState('');
	const [password, setPassword] = useState('');
	const [showPassword, setShowPassword] = useState(false);
	const navigate = useNavigate();
	const handleLogin = () => {

		fetchUserData(email, password)
			.then(data => {
				props.setIsLoggedIn(true);
				props.setEmail(data.data.email);
				props.setID(data.data.id);
				localStorage.setItem('session_token', data.data.session_token);

				navigate('/');
			})
			.catch(error => {
				console.error(error)
				alert('Неверные учетные данные. Пожалуйста, попробуйте еще раз.');
			});
	};

	return (
		<div className="login-container">
			<div className="login-box">
				<div className="center">
					<h2 style={{ color: '#FBFF33', fontWeight: 'bold'}}>Вход</h2>
					<p>Добро пожаловать! Пожалуйста, войдите в систему, чтобы получить доступ к своей учетной записи.</p>
				</div>
				<div className="input-group">
					<input
						type="text"
						value={email}
						onChange={(e) => setEmail(e.target.value)}
						placeholder="Введите Логин"
						className="login"
					/>
					<input
						type={showPassword ? 'text' : 'password'}
						value={password}
						onChange={(e) => setPassword(e.target.value)}
						placeholder="Введите Пароль"
						className="password"
					/>
					<label className="eye-icon" onClick={() => setShowPassword(!showPassword)}>
						<img
							src={showPassword ? EyeOpened : EyeClosed}
							alt={showPassword ? 'Open Eye' : 'Closed Eye'}
						/>
					</label>
				</div>
				<div className="center">
					<button onClick={handleLogin}>Войти</button>
				</div>
			</div>
		</div>
	);
}

export default Login;
