// eslint-disable-next-line no-unused-vars
import React from 'react';
import './Home.css';
import { useNavigate } from 'react-router-dom';
import video from './assets/aboutUs.mp4';
import aboutUsIcon from './assets/about-us-icon.png';
import arrow from './assets/arrow.svg';

const videoStyle = {
	width: '100%',
	height: 'auto',
	border: '1px solid #1A1A1A',
	borderRadius: '10px',
	boxShadow: '0 0 10px rgba(0, 0, 0, 0.1)'
};

function Home() {
	const navigate = useNavigate();
	const handleLoginClick = () => {
		navigate('/login');
	};

	const servicePoints = [
		'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
		'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
		'Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
		'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
		'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.',
		'Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
	];

	return (
		<div className="content">
			<div className="home-container">
				<div className="content-left">
					<h1>
						<span className="content-title" style={{ color: '#FBFF33', fontWeight: 'bold'}}>CallChanSolutions</span> - лучшее решение для вашей компании.
					</h1>
					<p className="content-text">
						Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
						dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
						commodo consequat.
					</p>
					<button className="content-button" onClick={handleLoginClick}>
						Вход
					</button>
				</div>
				<div className="content-right">
					<video src={video} controls style={videoStyle} />
				</div>
			</div>
			<div className="service">
				<h2>
					О нашем <span className="service-title" style={{ color: '#FBFF33', fontWeight: 'bold'}}>сервисе</span>
				</h2>
				<div className="service-container">
					<img src={aboutUsIcon} alt="about us" className="service-icon" />
					<div className="service-content">
						<h3>Lorem Lorem:</h3>
						<ul className="service-list">
							{servicePoints.map((point, index) => (
								<li key={index}>{point}</li>
							))}
						</ul>
					</div>
				</div>
			</div>
			<div className="advantages-container">
				<h2>Преимущества нашего <span style={{ color: '#FBFF33', fontWeight: 'bold'}}>сервиса</span></h2>
				<div className="row">
					<div className="advantage">
						<div className="advantage-header">
							<h3>Безопасность</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
							dolore magna aliqua.</p>
					</div>
					<div className="advantage advantage-right">
						<div className="advantage-header">
							<h3>Универсальность</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
							dolore magna aliqua.</p>
					</div>
				</div>
				<div className="row">
					<div className="advantage">
						<div className="advantage-header">
							<h3>Быстрый доступ</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
							dolore magna aliqua.</p>
					</div>
					<div className="advantage advantage-right">
						<div className="advantage-header">
							<h3>Надежность</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et
							dolore magna aliqua.</p>
					</div>
				</div>
			</div>
		</div>
	);
}

export default Home;
