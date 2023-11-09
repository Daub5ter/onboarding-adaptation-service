import { useState, useRef } from "react";
import {debugLog} from "express-fileupload/lib/utilities.js";

const hostUrl = "/upload";

export const UploadFile = () => {
    const filePicker = useRef(null);
    const [selectedFile, setSelectedFile] = useState(null);
    const [uploaded, setUploaded] = useState();

    const handleChange = (event) => {
        console.log(event.target.files);
        setSelectedFile(event.target.files[0]);
    };

    const handleUpload = async () => {
        if (!selectedFile) {
            alert("Please select a file");
            return;
        }

        const formData = new FormData();
        formData.append('file', selectedFile);

        const res = await fetch(hostUrl, {
            method: 'POST',
            body: formData,
        });
        const data = await res.json();

        setUploaded(data);
    };

    const handlePick = () => {
        filePicker.current.click();
    }

    return (
        <>
            <button onClick={handlePick}>Pick file</button>
            <input
                className="hidden"
                type="file"
                ref = {filePicker}
                onChange={handleChange}
                accept="image/*,.png,.jpg,.jpeg,.gif,.web"
            />

            <button onClick={handleUpload}>Upload now</button>

            {selectedFile && (
                <ul>
                    <li>Name: {selectedFile.name}</li>
                    <li>Type: {selectedFile.type}</li>
                    <li>Size: {selectedFile.size}</li>
                </ul>
            )}

            {uploaded && (
                <div>
                    <h2>{uploaded.fileName}</h2>
                    <img alt='' src={uploaded.filePath} width="200"/>
                </div>
            )}
        </>
    );
};