// eslint-disable-next-line no-unused-vars
import React, { useState } from 'react';
import './Onboarding.css';
import uncheckCircle from './assets/uncheck-circle.svg';
import checkCircle from './assets/check-circle.svg';

function Onboarding() {

	const [statuses, setStatuses] = useState([
		{ title: 'Должность 1', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et\n' +
				'\t\t\t\t\t\tdolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea\n' +
				'\t\t\t\t\t\tcommodo consequat.', checked: false },
		{ title: 'Должность 2', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et\n' +
				'\t\t\t\t\t\tdolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea\n' +
				'\t\t\t\t\t\tcommodo consequat.', checked: false },
		{ title: 'Должность 3', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et\n' +
				'\t\t\t\t\t\tdolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea\n' +
				'\t\t\t\t\t\tcommodo consequat.', checked: false },
		{ title: 'Должность 4', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et\n' +
				'\t\t\t\t\t\tdolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea\n' +
				'\t\t\t\t\t\tcommodo consequat.', checked: false }
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
