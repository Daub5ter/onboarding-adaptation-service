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
		'Наша компания предлагает услуги лёгкого онбординга и адаптации сотрудников.',
		'Мы поможем вашим новым сотрудникам быстро и комфортно войти в рабочий процесс, познакомим их с коллективом и корпоративной культурой, а также обеспечим необходимую поддержку на первых этапах работы.',
		'Наша цель - создать благоприятную атмосферу для успешного старта новых сотрудников и увеличить их мотивацию и эффективность в работе.',
		'Доверьте нам заботу о ваших сотрудниках, и вы увидите, как они станут неотъемлемой частью вашей команды!',
	];

	return (
		<div className="content">
			<div className="home-container">
				<div className="content-left">
					<h1>
						<span className="content-title" style={{ color: '#FBFF33', fontWeight: 'bold'}}>CallChanSolutions</span> - лучшее решение для вашей компании.
					</h1>
					<p className="content-text">
						С нами адаптация сотрудников станет легкой и приятной. Опробуйте наш сервис и свяжитесь с нами
						Решение трека "Сервис онбординга и адаптации сотрудников"
						Московский инновационный кластер. Лидеры цифровой трансформации & PROSCOM
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
						<h3>Почему именно мы?</h3>
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
						<p>Мы используем современные технологии быстрой и безопасной передачи данных,
							легко поддерживая сервис.</p>
					</div>
					<div className="advantage advantage-right">
						<div className="advantage-header">
							<h3>Универсальность</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Наша компания предлагает интерактивно простые и понятные инструкции для легкого понимая
							своей работы сотрудником и решение различных возникающих вопросов и проблем.</p>
					</div>
				</div>
				<div className="row">
					<div className="advantage">
						<div className="advantage-header">
							<h3>Быстрый доступ</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Мы работаем непосредственно с Вами, предоставляя всевозможные функции, взависмости от желаний
							и нужд.</p>
					</div>
					<div className="advantage advantage-right">
						<div className="advantage-header">
							<h3>Надежность</h3>
							<img src={arrow} alt="Стрелка" />
						</div>
						<p>Мы предоставляем возможность создавать собственные информационные блоки и инструкции.
							Реализуем интерактивные видео для полного погружения сотрудника в новую среду.</p>
					</div>
				</div>
			</div>
		</div>
	);
}

export default Home;
