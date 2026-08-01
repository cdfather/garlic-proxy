import express from 'express';
import { GoogleGenAI } from '@google/genai';

const app = express();
app.use(express.json());

const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

// Test için kök dizin
app.get('/', (req, res) => {
    res.send("GarlicAI Proxy Server is Online!");
});

// Terminalden gelen isteklerin işlendiği kısım
app.post('/api/chat', async (req, res) => {
    try {
        const { prompt } = req.body;
        if (!prompt) {
            return res.status(400).json({ error: "Prompt is required." });
        }

        const response = await ai.models.generateContent({
            model: 'gemini-2.5-flash',
            contents: prompt,
            config: {
                systemInstruction: "You are GarlicAI, a fast and helpful assistant running inside a Linux terminal. Provide clear, direct, concise, and accurate answers in English."
            }
        });

        res.json({ result: response.text });
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: "Internal server error." });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
