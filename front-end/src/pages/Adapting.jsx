// eslint-disable-next-line no-unused-vars
import React, {useEffect, useState} from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';
import './Adapting.css';
import PrevArrow from './assets/arrow-left.svg';
import NextArrow from './assets/arrow-right.svg';
import img1 from './assets/1Cfirst.png';
import img4 from './assets/1Cthird.png';
import img2 from './assets/1Cfouth.png'
import img5 from './assets/sysadmin1.png'
import {useNavigate} from "react-router-dom";

const slides = [
	[
		{
			description: 'Подключение через SSH к серверу в Linux',
			mediaUrl: img1
		},
		{
			description: 'Раздел:Отчеты–Экспресс-проверка.'
		},
	],
	[
		{
			description: 'Укажите период формирования отчета и организацию, по которой проводится экспресс-проверка. Для получения достоверной информации при формировании отчета по разделам, не связанным с учетом НДС, выбирайте периоды, кратные календарному месяцу.',
			mediaUrl: img2
		},
	],
	[
		{
			description: 'Если пытаетесь подключится через SSH к этому серверу первый раз, то утилита также попросит подтвердить добавление нового устройства в свой список известных устройств, здесь нужно набрать yes и нажать Enter.',
			mediaUrl: img5
		},
		{
			description: 'Теперь вы подключены, и все вводимые далее команды будут выполнены на удаленном сервере.'
		},
	],
	[
		{
			description: 'Детализируйте отчет, раскрыв соответствующие подразделы по знаку "+". В отчет выводится подробная информация о результатах проверки, возможных причинах возникновения ошибок, рекомендациях по их устранению. По некоторым подразделам отчета возможна детализация вплоть до первичного документа, который может быть открыт для редактирования или просмотра по двойному щелчку мыши на соответствующей строке отчета.',
			mediaUrl: img4
		},
		{
			description: 'Следуя рекомендациям отчета, устраните выявленные ошибки и повторите экспресс-проверку. Рекомендации отчета можно проигнорировать, если есть уверенность, что проблемная ситуация не является ошибкой и (или) не повлияет на правильность ведения учета и формирования отчетности.'
		},
	],
];

function Adapting(props) {
	const navigate = useNavigate();

	const handleNavigate = () => {
		useEffect(() => {
			navigate('/login');
		}, []);
	}

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
		<> {props.isLoaded && props.isLoggedIn ?
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
								<h3>Подключение к серверу через SSH в Linux</h3>
								<ul className="instruction-list">
									{stepList.map((item, stepIndex) => (
										<li className="instruction-step" key={stepIndex}>
											<h4></h4>
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
			: handleNavigate()}
		</>
	);
}

export default Adapting;
