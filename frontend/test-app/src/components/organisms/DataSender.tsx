import React, { useState } from 'react';
import axios from 'axios';

const DataSender: React.FC = () => {
    const [inputValue, setInputValue] = useState('');

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setInputValue(e.target.value);
    };

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/api/data', {
                message: inputValue
            });
            console.log('Server Response:', response.data);
        } catch (error) {
            console.error('Error sending data:', error);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input type="text" value={inputValue} onChange={handleInputChange} />
            <button type="submit">Send</button>
        </form>
    );
};

export default DataSender;

