import React, { useEffect, useState } from 'react';
import axios from 'axios';

const DataDisplay = () => {
    const [data, setData] = useState(null);

    useEffect(() => {
        // Запрос к API
        axios.get('http://localhost:8080/api/your-endpoint')
            .then(response => {
                setData(response.data);
            })
            .catch(error => {
                console.error('Ошибка при получении данных:', error);
            });
    }, []);

    return (
        <div>
            <h1>Данные с API</h1>
            {data ? (
                <pre>{JSON.stringify(data, null, 2)}</pre>
            ) : (
                <p>Загрузка данных...</p>
            )}
        </div>
    );
};

export default DataDisplay;
