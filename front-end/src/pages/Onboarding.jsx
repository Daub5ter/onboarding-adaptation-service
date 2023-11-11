// eslint-disable-next-line no-unused-vars
import React, { useState } from 'react';
import './Onboarding.css';
import uncheckCircle from './assets/uncheck-circle.svg';
import checkCircle from './assets/check-circle.svg';

function Onboarding() {

	const [statuses, setStatuses] = useState([
		{ title: 'Компания', description: 'Test Company - это инновационная IT-компания, специализирующаяся на предоставлении решений в области информационных технологий. Мы предлагаем широкий спектр услуг, включая разработку программного обеспечения, веб-разработку, мобильные приложения, облачные решения, консалтинг и IT-аутсорсинг.\n'+
				'Мы стремимся помочь нашим клиентам достичь цифровой трансформации и повысить их эффективность, конкурентоспособность и инновационность. Наша команда высококвалифицированных специалистов имеет глубокие знания и опыт работы в различных отраслях, что позволяет нам создавать индивидуальные IT-решения, наиболее соответствующие потребностям каждого клиента.\n' +
				'\n' +
				'Мы следим за последними технологическими тенденциями и интегрируем их в наши проекты, чтобы обеспечить передовые решения и удовлетворить ожидания наших клиентов. Мы ценим инновации, эффективность и качество во всем, что делаем.\n' , checked: false },
		{ title: 'Отделы', description: 'Отдел разработки программного обеспечения:\n' +
				'Отдел разработки программного обеспечения в Test Company состоит из высококвалифицированных разработчиков, специализирующихся на создании инновационных и надежных программных решений. Они работают в тесном сотрудничестве с клиентами, чтобы полностью понять их требования и создать индивидуальные продукты, соответствующие их потребностям.\n' +
				'Отдел веб-разработки:\n' +
				'Отдел веб-разработки занимается созданием функциональных и эстетически привлекательных веб-сайтов и приложений. Наши веб-разработчики обладают глубокими знаниями веб-технологий и следят за последними трендами дизайна и пользовательского опыта, чтобы создавать уникальные и интуитивно понятные пользовательские интерфейсы.\n' +
				'\n' +
				'Отдел мобильной разработки:\n' +
				'Отдел мобильной разработки занимается созданием инновационных мобильных приложений для различных платформ, включая iOS и Android. Наши разработчики мобильных приложений имеют глубокие знания в области мобильных технологий и следят за последними трендами в разработке приложений для обеспечения передовых решений.\n' +
				'\n' +
				'Отдел облачных решений:\n' +
				'Отдел облачных решений в Test Company  предоставляет услуги по разработке и внедрению облачных архитектур и решений. Наши специалисты в области облачных технологий помогают клиентам оптимизировать и масштабировать свои бизнес-процессы, обеспечивая гибкость, безопасность и доступность данных.\n' +
				'Отдел консалтинга:\n' +
				'Отдел консалтинга предоставляет экспертные консультационные услуги по вопросам информационных технологий, стратегического планирования и цифровой трансформации. Наши консультанты работают в тесном сотрудничестве с клиентами, помогая им определить и реализовать оптимальные IT-решения для достижения их бизнес-целей.\n' +
				'\n' +
				'Отдел IT-аутсорсинга:\n' +
				'Отдел IT-аутсорсинга предоставляет услуги по полной или частичной передаче функций IT-управления и поддержки. Мы предлагаем гибкие решения, адаптированные к потребностям клиентов, чтобы обеспечить надежное функционирование IT-инфраструктуры и оптимизировать затраты на IT-обслуживание.\n' +
				'\n' +
				'Каждый отдел в Test Company  является частью нашей команды, работающей на достижение общих целей компании и удовлетворение потребностей наших клиентов. Мы гордимся нашей экспертизой, профессионализмом и инновационным подходом, который позволяет нам предоставлять высококачественные IT-решения.\n', checked: false },
		{ title: 'Дресс-код', description: 'Дресс-код компании "Test Company" отражает нашу профессиональную и деловую атмосферу, а также учитывает комфорт сотрудников. Мы ставим на первое место профессионализм и представительность, однако не жестко ограничиваем стиль одежды, чтобы сотрудники могли выразить свою индивидуальность. Вот общие рекомендации по дресс-коду в нашей компании:\n' +
				'\n' +
				'Бизнес-формальный:\n' +
				'В рамках бизнес-формального дресс-кода ожидается профессиональный и официальный вид. Мужчины могут носить костюмы или брючные костюмы с рубашками и галстуками. Женщины могут выбирать между костюмами, платьями или юбками с блузками. Рекомендуется носить закрытую обувь и аккуратные аксессуары.\n' +
				'\n' +
				'Бизнес-кэжуал:\n' +
				'В рамках бизнес-кэжуал дресс-кода можно носить менее формальную одежду, сохраняя при этом профессиональный вид. Мужчины могут носить брюки или хорошо подходящие джинсы в сочетании с рубашками или поло. Женщины могут выбирать между брюками, юбками или платьями, допускается носить блузки или свитеры. Обувь должна быть аккуратной, а аксессуары - умеренными.\n' +
				'\n' +
				'Кежуал:\n' +
				'В рамках кежуал дресс-кода допускается более неформальный стиль одежды. Мужчины и женщины могут носить джинсы, худи, футболки или поло, однако они все равно должны выглядеть аккуратно и профессионально. Обувь и аксессуары также должны быть ухоженными и подходящими.\n' +
				'\n' +
				'Важно помнить, что в некоторых ситуациях, например, при встрече с клиентами или на корпоративных мероприятиях, может потребоваться более формальный дресс-код. В таких случаях сотрудники получают соответствующую информацию и инструкции.\n' +
				'\n' +
				'Наша цель - создать профессиональную и уважительную атмосферу, поэтому мы рекомендуем сотрудникам следовать общим принципам дресс-кода и поддерживать нашу корпоративную имиджевую политику.', checked: false },
		{ title: 'Проекты', description: 'Компания "Test Company" занимается разработкой следующих проектов в области backend, frontend, UE5 (Unreal Engine 5) и Blender:\n' +
				'\n' +
				'Проект backend-разработки:\n' +
				'Ответственный: Цебро Максим, руководитель отдела backend-разработки.\n' +
				'Описание проекта: Команда backend-разработчиков занимается созданием серверной части программного обеспечения. Они отвечают за разработку и поддержку баз данных, веб-серверов, API и других компонентов, необходимых для обработки данных и взаимодействия с фронтендом. Цель проекта - обеспечить стабильную и эффективную работу серверной инфраструктуры.\n' +
				'\n' +
				'Проект frontend-разработки:\n' +
				'Ответственный: Елизавета Криворучко, руководитель отдела frontend-разработки.\n' +
				'Описание проекта: Команда frontend-разработчиков занимается созданием пользовательского интерфейса и клиентской части программного обеспечения. Они отвечают за разработку и поддержку веб-интерфейсов, мобильных приложений и других пользовательских приложений. Цель проекта - предоставить пользователю удобный и интуитивно понятный интерфейс для взаимодействия с системой.\n' +
				'\n' +
				'Проект разработки на Unreal Engine 5:\n' +
				'Ответственный: Эрнес Самадинов, руководитель отдела UE5-разработки.\n' +
				'Описание проекта: Команда разработчиков на Unreal Engine 5 занимается созданием игровых и визуализационных проектов, используя передовые возможности этого игрового движка. Они отвечают за создание игровых механик, моделей, текстур, освещения и других компонентов, необходимых для реализации проектов на UE5. Цель проекта - создать увлекательные и реалистичные визуальные проекты.\n' +
				'\n' +
				'Проект разработки на Blender:\n' +
				'Ответственный: Семен Святаш, руководитель отдела Blender-разработки.\n' +
				'Описание проекта: Команда разработчиков на Blender занимается созданием 3D-моделей, анимаций и визуализаций. Они отвечают за моделирование, текстурирование, анимацию и рендеринг объектов и сцен с использованием функциональности Blender. Цель проекта - создание высококачественных и креативных 3D-контента.\n' +
				'\n' +
				'Каждый проект имеет своего ответственного руководителя, который отвечает за планирование, координацию и успешное выполнение проекта. Они работают в тесном сотрудничестве с другими отделами компании, чтобы достичь поставленных целей и обеспечить высокое качество реализации проектов.', checked: false }
	]);

	const [counter, setCounter] = useState(0);

	const toggleStatus = (index) => {
		const newStatuses = [...statuses];
		newStatuses[index].checked = !newStatuses[index].checked;
		setStatuses(newStatuses);

		if (newStatuses[index].checked) {
			setCounter((prevCounter) => prevCounter + 1);
		} else {
			setCounter((prevCounter) => prevCounter - 1);
		}
	};

	const isAllRead = statuses.every(item => item.checked);

	return (
		<>
			<div className="onboarding">
				<h2>Материал к ознакомлению:</h2>
				{isAllRead ? (
					<p className="onboarding-counter">Вы ознакомились со всеми пунктами! </p>
				) : (
					<p className="onboarding-counter">Количество пунктов, с которыми вы ознакомились: {counter} </p>
				)}
				{statuses.map((item, index) => (
					<div className="onboarding-container" key={index}>
						<div className="onboarding-left">
							<h3>{item.title}</h3>
							<p>{item.description}</p>
						</div>
						<img
							src={item.checked ? checkCircle : uncheckCircle}
							alt="Стрелка"
							onClick={() => toggleStatus(index)}
						/>
					</div>
				))}
			</div>
		</>
	);
}

export default Onboarding;