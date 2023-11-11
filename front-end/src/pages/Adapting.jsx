// eslint-disable-next-line no-unused-vars
import React, { useState } from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import './Adapting.css';
import PrevArrow from './assets/arrow-left.svg';
import NextArrow from './assets/arrow-right.svg';


const slides = [
	[
		{
			title: '',
			description: 'Отчет "Экспресс-проверка ведения учета" позволяет провести быструю проверку правильности отражения операций в программе, в том числе проанализировать состояние бухгалтерского учета.',
			mediaUrl: 'https://play-lh.googleusercontent.com/HUuQc4Zpl6x7fUyX-jFMmcuUbW77REw4UKl5rfhHfP4VY6ctBU1w1I_RZWsXaojFgIo=w480-h960-rw'
		},
		{
			title: 'Шаг 1',
			description: 'Раздел:Отчеты–Экспресс-проверка.'
		},
	],
	[
		{
			title: 'Шаг 4',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore',
			mediaUrl: 'https://goods-photos.static1-sima-land.com/items/4668517/0/1600.jpg?v=1605611154'
		},
		{
			title: 'Шаг 5',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing'
		},
		{
			title: 'Шаг 6',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod ididunt'
		}
	],
	[
		{
			title: 'Шаг 4',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore',
			mediaUrl: 'https://goods-photos.static1-sima-land.com/items/4668517/0/1600.jpg?v=1605611154'
		},
		{
			title: 'Шаг 5',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing'
		},
		{
			title: 'Шаг 6',
			description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod ididunt'
		}
	],
];

function Adapting() {
	const [currentSlide, setCurrentSlide] = useState(0);

	const handlePrevClick = () => {
		setCurrentSlide((prevSlide) => (prevSlide === 0 ? slides.length - 1 : prevSlide - 1));
	};

	const handleNextClick = () => {
		setCurrentSlide((prevSlide) => (prevSlide === slides.length - 1 ? 0 : prevSlide + 1));
	};

	const sliderSettings = {
		dots: true,
		infinite: true,
		speed: 500,
		slidesToShow: 1,
		slidesToScroll: 1,
		prevArrow: <img src={PrevArrow} alt="Previous" onClick={handlePrevClick} />,
		nextArrow: <img src={NextArrow} alt="Next" onClick={handleNextClick} />
	};

	return (
		<div className="adapting-container">
			<h2 style={{ color: '#FBFF33', fontWeight: 'bold' }}>Инструкция</h2>
			<Slider {...sliderSettings}>
				{slides.map((stepList, index) => (
					<div key={index}>
						<div className="adapting-content">
							<div className="photo-column">
								<img src={stepList[currentSlide].mediaUrl} alt="to do" />
							</div>
							<div className="instruction-column">
								<h3>Как в "1С:Бухгалтерии" выполнить экспресс-проверку состояния бухгалтерского учета?</h3>
								<ul className="instruction-list">
									{stepList.map((item, stepIndex) => (
										<li className="instruction-step" key={stepIndex}>
											<h4>{item.title}</h4>
											<p>{item.description}</p>
										</li>
									))}
								</ul>
							</div>
						</div>
					</div>
				))}
			</Slider>
		</div>
	);
}

export default Adapting;
