document.getElementById('shorten-form').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const url = document.getElementById('url-input').value;
    const customName = document.getElementById('custom-input').value;
    
    try {
        const response = await fetch('/api/shorten', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ url, custom_name: customName })
        });
        
        const data = await response.json();
        
        // Показываем результат
        const shortLink = document.getElementById('short-link');
        shortLink.href = `/s/${data.short_url}`;
        shortLink.textContent = `${window.location.origin}/s/${data.short_url}`;
        
        document.getElementById('result').style.display = 'block';
        
        // Кнопка аналитики
        document.getElementById('analytics-btn').onclick = () => {
            window.open(`/ui/analytics/${data.short_url}`, '_blank');
        };
        
    } catch (error) {
        alert('Error creating short URL');
    }
});