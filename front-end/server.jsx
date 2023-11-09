const express = require('express');
const fileUpload = require('express-fileupload');

const app = express();

app.use(fileUpload({
    createParentPath: true,
}))

app.post('/upload', (req, res) => {
    if (!req.files) {
        return res.status(400).json({ msg: 'No file uploaded' });
    }

    const file = req.files.file;

    if (!file) return res.json({ error: 'Incorrect input name' });

    file.mv(`${__dirname}/public/uploads/${file.name}`, err => {
        if (err) {
            console.error(err);
            return res.status(500).send(err);
        }
        console.log('file was uploaded');

        res.json({
            fileName: file.name,
            filePath: `/public/uploads/${file.name}`,
        });
    });
});

app.listen(1234, () => console.log('Server Started...'));