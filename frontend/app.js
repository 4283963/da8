const API_BASE = '/api';

document.addEventListener('DOMContentLoaded', function() {
    bindSliderEvents();
    bindFormSubmit();
    bindRefreshButton();
    loadBraceletList();
});

function bindSliderEvents() {
    const transRange = document.getElementById('translucencyRange');
    const transInput = document.getElementById('translucency');
    const fineRange = document.getElementById('finenessRange');
    const fineInput = document.getElementById('fineness');

    transRange.addEventListener('input', function() {
        transInput.value = this.value;
    });

    transInput.addEventListener('input', function() {
        transRange.value = this.value || 0;
    });

    fineRange.addEventListener('input', function() {
        fineInput.value = this.value;
    });

    fineInput.addEventListener('input', function() {
        fineRange.value = this.value || 0;
    });
}

function bindFormSubmit() {
    const form = document.getElementById('braceletForm');
    form.addEventListener('submit', async function(e) {
        e.preventDefault();

        const data = {
            name: document.getElementById('name').value.trim(),
            material: document.getElementById('material').value,
            translucency: parseFloat(document.getElementById('translucency').value),
            fineness: parseFloat(document.getElementById('fineness').value),
            bead_count: parseInt(document.getElementById('beadCount').value)
        };

        try {
            const response = await fetch(`${API_BASE}/bracelets`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();

            if (result.code === 200) {
                showResult(result.data);
                form.reset();
                document.getElementById('beadCount').value = 17;
                document.getElementById('translucencyRange').value = 50;
                document.getElementById('finenessRange').value = 50;
                loadBraceletList();
            } else {
                alert('提交失败：' + result.message);
            }
        } catch (error) {
            console.error('Error:', error);
            alert('网络错误，请稍后重试');
        }
    });
}

function bindRefreshButton() {
    const refreshBtn = document.getElementById('refreshBtn');
    refreshBtn.addEventListener('click', loadBraceletList);
}

function showResult(data) {
    const card = document.getElementById('resultCard');
    const scoreEl = document.getElementById('resultScore');
    const gradeEl = document.getElementById('resultGrade');

    scoreEl.textContent = data.score.toFixed(2);

    gradeEl.textContent = data.grade;
    gradeEl.className = 'grade-badge ' + data.grade;

    card.classList.remove('hidden');
}

async function loadBraceletList() {
    const tbody = document.getElementById('braceletList');
    tbody.innerHTML = '<tr><td colspan="9" class="loading">加载中...</td></tr>';

    try {
        const response = await fetch(`${API_BASE}/bracelets`);
        const result = await response.json();

        if (result.code === 200 && result.data && result.data.length > 0) {
            renderBraceletList(result.data);
        } else {
            tbody.innerHTML = '<tr><td colspan="9" class="loading">暂无记录</td></tr>';
        }
    } catch (error) {
        console.error('Error:', error);
        tbody.innerHTML = '<tr><td colspan="9" class="loading">加载失败，请刷新重试</td></tr>';
    }
}

function renderBraceletList(data) {
    const tbody = document.getElementById('braceletList');
    tbody.innerHTML = '';

    data.forEach(function(item) {
        const tr = document.createElement('tr');
        tr.innerHTML = `
            <td>${item.id}</td>
            <td>${escapeHtml(item.name)}</td>
            <td>${escapeHtml(item.material)}</td>
            <td>${item.translucency.toFixed(2)}</td>
            <td>${item.fineness.toFixed(2)}</td>
            <td>${item.bead_count}</td>
            <td><strong>${item.score.toFixed(2)}</strong></td>
            <td><span class="grade-badge ${escapeHtml(item.grade)}" style="padding: 2px 10px; font-size: 12px;">${escapeHtml(item.grade)}</span></td>
            <td><button class="delete-btn" data-id="${item.id}">删除</button></td>
        `;
        tbody.appendChild(tr);
    });

    tbody.querySelectorAll('.delete-btn').forEach(function(btn) {
        btn.addEventListener('click', function() {
            const id = this.getAttribute('data-id');
            deleteBracelet(id);
        });
    });
}

async function deleteBracelet(id) {
    if (!confirm('确定要删除这条记录吗？')) {
        return;
    }

    try {
        const response = await fetch(`${API_BASE}/bracelets/${id}`, {
            method: 'DELETE'
        });

        const result = await response.json();

        if (result.code === 200) {
            loadBraceletList();
        } else {
            alert('删除失败：' + result.message);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('网络错误，请稍后重试');
    }
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}
