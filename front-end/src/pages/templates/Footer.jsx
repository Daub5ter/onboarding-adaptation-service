// eslint-disable-next-line no-unused-vars
import React from 'react';
import './Footer.css';
import emailIcon from '../assets/mail.svg';
import phoneIcon from '../assets/phone.svg';
import addressIcon from '../assets/location.svg';

function Footer() {
	return (
		<footer className="footer">
			<div className="contact-info">
				<div className="contact-items">
					<img src={emailIcon} alt="Email Icon" className="icon" />
					<p>hello@mail.ru</p>
				</div>
				<div className="contact-items">
					<img src={phoneIcon} alt="Phone Icon" className="icon" />
					<p>8(800)555-35-35</p>
				</div>
				<div className="contact-items">
					<img src={addressIcon} alt="Address Icon" className="icon" />
					<p>Адрес компании</p>
				</div>
			</div>
			<div className="all-rights">
				<p>All rights reserved</p>
			</div>
		</footer>
	);
}

export default Footer;
