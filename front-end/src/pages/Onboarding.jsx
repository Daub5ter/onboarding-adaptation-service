// eslint-disable-next-line no-unused-vars
import React, {useEffect, useState} from 'react';
import './Onboarding.css';
import uncheckCircle from './assets/uncheck-circle.svg';
import checkCircle from './assets/check-circle.svg';
import {useNavigate} from "react-router-dom";
import LoadSession from "../Auth/LoadSession";

function fetchUserKnowledge(id) {
	const payload = {
		action: "get_all_knowledge",
		id: {
			id: id,
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

function fetchSolveKnowledge(userID, knowledgeID) {
	const payload = {
		action: "add_users_knowledge",
		users_known: {
			user_id: userID,
			knowledge_id: knowledgeID,
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

function fetchUserPercent(id) {
	const payload = {
		action: "get_percent_knowledge",
		id: {
			id: id,
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

function Onboarding(props) {
	const navigate = useNavigate();

	const handleNavigate = () => {
		useEffect(() => {
			navigate('/login');
		}, []);
	}

	const [statuses, setStatuses] = useState([])

	if (statuses.length === 0 && props.isLoggedIn) {
		fetchUserKnowledge(props.id)
			.then(data => {
				if (data.error !== true) {
					setStatuses(data.data);
				}
			})
			.catch(error => {
				console.error(error)
			});
	}

		/*const [statuses, setStatuses] = useState([
			{
				title: 'Компания',
				description: 'Test Company - это инновационная IT-компания, специализирующаяся на предоставлении решений в области информационных технологий. Мы предлагаем широкий спектр услуг, включая разработку программного обеспечения, веб-разработку, мобильные приложения, облачные решения, консалтинг и IT-аутсорсинг ' +
					'Мы стремимся помочь нашим клиентам достичь цифровой трансформации и повысить их эффективность, конкурентоспособность и инновационность. Наша команда высококвалифицированных специалистов имеет глубокие знания и опыт работы в различных отраслях, что позволяет нам создавать индивидуальные IT-решения, наиболее соответствующие потребностям каждого клиента. ' +
					'Мы следим за последними технологическими тенденциями и интегрируем их в наши проекты, чтобы обеспечить передовые решения и удовлетворить ожидания наших клиентов. Мы ценим инновации, эффективность и качество во всем, что делаем.',
				solved: false
			},
			{
				title: 'Отделы',
				description:
					'Отдел разработки программного обеспечения в Test Company состоит из высококвалифицированных разработчиков, специализирующихся на создании инновационных и надежных программных решений.' +
					'Отдел веб-разработки занимается созданием функциональных и эстетически привлекательных веб-сайтов и приложений.' +
					'Отдел мобильной разработки занимается созданием инновационных мобильных приложений для различных платформ, включая iOS и Android.' +
					'Отдел облачных решений в Test Company  предоставляет услуги по разработке и внедрению облачных архитектур и решений.' +
					'Отдел консалтинга предоставляет экспертные консультационные услуги по вопросам информационных технологий, стратегического планирования и цифровой трансформации.' +
					'Отдел IT-аутсорсинга предоставляет услуги по полной или частичной передаче функций IT-управления и поддержки.' +
					'Каждый отдел в Test Company  является частью нашей команды, работающей на достижение общих целей компании и удовлетворение потребностей наших клиентов.',
				solved: false
			},
			{
				title: 'Дресс-код',
				description: 'Дресс-код компании "Test Company" отражает нашу профессиональную и деловую атмосферу, а также учитывает комфорт сотрудников. Мы ставим на первое место профессионализм и представительность, однако не жестко ограничиваем стиль одежды, чтобы сотрудники могли выразить свою индивидуальность. Вот общие рекомендации по дресс-коду в нашей компании:' +
					'Бизнес-формальный:' +
					'В рамках бизнес-формального дресс-кода ожидается профессиональный и официальный вид. Мужчины могут носить костюмы или брючные костюмы с рубашками и галстуками. Женщины могут выбирать между костюмами, платьями или юбками с блузками. Рекомендуется носить закрытую обувь и аккуратные аксессуары.' +
					'Наша цель - создать профессиональную и уважительную атмосферу, поэтому мы рекомендуем сотрудникам следовать общим принципам дресс-кода и поддерживать нашу корпоративную имиджевую политику.',
				solved: false
			},
			{
				title: 'Проекты',
				description: 'Компания "Test Company" занимается разработкой следующих проектов в области backend, frontend, UE5 (Unreal Engine 5) и Blender:\n' +
					'Проект backend-разработки: ' +
					'Ответственный: Цебро Максим, руководитель отдела backend-разработки.' +
					'Описание проекта: Команда backend-разработчиков занимается созданием серверной части программного обеспечения.' +
					'Проект frontend-разработки: ' +
					'Ответственный: Елизавета Криворучко, руководитель отдела frontend-разработки.' +
					'Описание проекта: Команда frontend-разработчиков занимается созданием пользовательского интерфейса и клиентской части программного обеспечения.' +
					'Проект разработки на Unreal Engine 5: ' +
					'Ответственный: Эрнес Самадинов, руководитель отдела UE5-разработки.' +
					'Описание проекта: Команда разработчиков на Unreal Engine 5 занимается созданием игровых и визуализационных проектов, используя передовые возможности этого игрового движка.' +
					'Проект разработки на Blender: ' +
					'Ответственный: Семен Святаш, руководитель отдела Blender-разработки.' +
					'Описание проекта: Команда разработчиков на Blender занимается созданием 3D-моделей, анимаций и визуализаций.' +
					'Каждый проект имеет своего ответственного руководителя, который отвечает за планирование, координацию и успешное выполнение проекта.',
				solved: false
			}
		]);*/

	const [percent, setPercent] = useState(0);
	const [isPercentGot, getPercent] = useState(false);

	if (!isPercentGot && props.isLoggedIn) {
		fetchUserPercent(props.id)
			.then(data => {
				if (data.error !== true) {
					setPercent(data.data);
					getPercent(true);
				}
			})
			.catch(error => {
				console.error(error)
			});
	}

	const toggleStatus = (index) => {
		const newStatuses = [...statuses];

		fetchSolveKnowledge(props.id, newStatuses[index].id)
			.then(dataSolve => {
				if (dataSolve.error !== true) {
					newStatuses[index].solved = true;
					setStatuses(newStatuses);

					fetchUserPercent(props.id)
						.then(data => {
							if (data.error !== true) {
								setPercent(data.data);
							}
						})
						.catch(error => {
							console.error(error)
						});
				}
			})
			.catch(error => {
				console.error(error)
			});
	};

	const isAllRead = statuses.every(item => item.solved);

	return (
			<> {props.isLoaded && props.isLoggedIn ?
					<div className="onboarding">
						<h2>Материал к ознакомлению:</h2>
						{isAllRead ? (
							<p className="onboarding-percent">Вы ознакомились со всеми пунктами! </p>
						) : (
							<p className="onboarding-percent">Процент пунктов, с которыми вы ознакомились: {percent}% </p>
						)}
						{statuses.map((item, index) => (
							<div className="onboarding-container" key={index}>
								<div className="onboarding-left">
									<h3>{item.title}</h3>
									<p>{item.description}</p>
								</div>
								<img
									src={item.solved ? checkCircle : uncheckCircle}
									alt="Стрелка"
									onClick={() => toggleStatus(index)}
								/>
							</div>
						))}
					</div>
					: handleNavigate() }
				</>
		);
}
export default Onboarding;
