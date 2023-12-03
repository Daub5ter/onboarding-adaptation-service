// eslint-disable-next-line no-unused-vars
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import './Header.css';
import logo from '../assets/logo.svg';

function Header(props) {
	const { pathname } = useLocation();

	return (
		<div className="app-container">
			<header className="header">
				<div className="logo-container">
					<img className="logo-image" src={logo} alt="Лого" />
					<h1 className="logo-text">CallChanSolutions</h1>
				</div>
				<nav className="navigation">
					<Link className="nav-link" to="/" aria-current={pathname === '/' && 'page'}>
						Главная
					</Link>
					<Link className="nav-link" to="/onboarding" aria-current={pathname === '/onboarding' && 'page'}>
						Онбординг
					</Link>
					<Link className="nav-link" to="/adaptation" aria-current={pathname === '/adaptation' && 'page'}>
						Адаптация
					</Link>
				</nav>
				<Link className="login-link" to="/login">
					{props.isLoggedIn ? <a>{props.username}</a> : <a>Вход</a>}
				</Link>
			</header>
		</div>
	);
}

export default Header;
