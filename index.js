import express from 'express';
import { GoogleGenAI } from '@google/genai';

const app = express();
app.use(express.json());

const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

app.post('/api/chat', async (req, res) => {
    try {
        const { prompt } = req.body;
        if (!prompt) {
            return res.status(400).json({ error: "Prompt gerekli." });
        }

        const response = await ai.models.generateContent({
            model: 'gemini-2.5-flash',
            contents: prompt,
            config: {
                systemInstruction: "Sen Linux terminalinde çalışan GarlicAI asistansın. Kısa, net ve yapıcı Türkçe cevaplar ver."
            }
        });

        res.json({ result: response.text });
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: "Sunucu hatası oluştu." });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Sunucu ${PORT} portunda aktif.`));
