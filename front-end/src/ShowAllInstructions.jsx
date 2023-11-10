import ReactDOM from "react";

export const ShowAllInstructions = () => {
    const items = [1, 2, 3, 4];

    ReactDOM.render(
        (<div>{items.map(item => <div>{item}</div>)}</div>),
        document.getElementById('container')
    );

}