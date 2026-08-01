import express from 'express';
import { GoogleGenAI } from '@google/genai';

const app = express();
app.use(express.json());

// Environment Variable üzerinden API Key alınıyor
const ai = new GoogleGenAI({ apiKey: process.env.GEMINI_API_KEY });

app.get('/', (req, res) => {
    res.send("GarlicAI Proxy Server is Online!");
});

app.post('/api/chat', async (req, res) => {
    try {
        const { prompt } = req.body;
        if (!prompt) {
            return res.status(400).json({ error: "Prompt is required." });
        }

        // Gemini API tarafında garantili çalışan model ismi
        const response = await ai.models.generateContent({
            model: 'gemini-1.5-flash',
            contents: prompt,
            config: {
                systemInstruction: "You are GarlicAI, a fast and helpful assistant running inside a Linux terminal. Provide clear, direct, concise, and accurate answers in English."
            }
        });

        res.json({ result: response.text });
    } catch (error) {
        console.error("Gemini API Error:", error);
        res.status(500).json({ 
            error: error.message || "Internal server error." 
        });
    }
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
