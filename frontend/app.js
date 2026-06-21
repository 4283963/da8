const API_BASE = '/api';
let currentDetailId = null;

document.addEventListener('DOMContentLoaded', function() {
    bindSliderEvents();
    bindFormSubmit();
    bindRefreshButton();
    bindDetailPanelEvents();
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
            fineness: parseFloat(document.getElementById('fineness').value)
        };

        const beadCountStr = document.getElementById('beadCount').value;
        if (beadCountStr !== '' && beadCountStr !== null) {
            const bc = parseInt(beadCountStr);
            if (!isNaN(bc) && bc >= 1) {
                data.bead_count = bc;
            }
        }

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
                document.getElementById('beadCount').value = '';
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
        const beadDisplay = item.bead_count === null || item.bead_count === undefined ? '—' : item.bead_count;
        tr.innerHTML = `
            <td>${item.id}</td>
            <td>${escapeHtml(item.name)}</td>
            <td>${escapeHtml(item.material)}</td>
            <td>${item.translucency.toFixed(2)}</td>
            <td>${item.fineness.toFixed(2)}</td>
            <td>${beadDisplay}</td>
            <td><strong>${item.score.toFixed(2)}</strong></td>
            <td><span class="grade-badge ${escapeHtml(item.grade)}" style="padding: 2px 10px; font-size: 12px;">${escapeHtml(item.grade)}</span></td>
            <td>
                <button class="view-btn" data-id="${item.id}">详情</button>
                <button class="delete-btn" data-id="${item.id}">删除</button>
            </td>
        `;
        tbody.appendChild(tr);
    });

    tbody.querySelectorAll('.delete-btn').forEach(function(btn) {
        btn.addEventListener('click', function() {
            const id = this.getAttribute('data-id');
            deleteBracelet(id);
        });
    });

    tbody.querySelectorAll('.view-btn').forEach(function(btn) {
        btn.addEventListener('click', function() {
            const id = this.getAttribute('data-id');
            viewBraceletDetail(id);
        });
    });
}

function bindDetailPanelEvents() {
    const closeBtn = document.getElementById('closeDetailBtn');
    closeBtn.addEventListener('click', function() {
        document.getElementById('detailPanel').classList.add('hidden');
        currentDetailId = null;
    });

    const generateBtn = document.getElementById('generateCardBtn');
    generateBtn.addEventListener('click', function() {
        if (currentDetailId) {
            downloadCard(currentDetailId);
        }
    });
}

async function viewBraceletDetail(id) {
    try {
        const response = await fetch(`${API_BASE}/bracelets/${id}`);
        const result = await response.json();

        if (result.code === 200 && result.data) {
            currentDetailId = id;
            renderDetailPanel(result.data);
            document.getElementById('detailPanel').classList.remove('hidden');
            document.getElementById('detailPanel').scrollIntoView({ behavior: 'smooth', block: 'start' });
        } else {
            alert('加载详情失败：' + result.message);
        }
    } catch (error) {
        console.error('Error:', error);
        alert('网络错误，请稍后重试');
    }
}

function renderDetailPanel(item) {
    const content = document.getElementById('detailContent');
    const beadDisplay = (item.bead_count === null || item.bead_count === undefined) ? '未填写' : item.bead_count + ' 颗';
    const createdAt = item.created_at ? new Date(item.created_at).toLocaleString('zh-CN') : '—';

    content.innerHTML = `
        <div class="detail-row"><span class="detail-label">手串编号</span><span class="detail-value">${escapeHtml(item.name)}</span></div>
        <div class="detail-row"><span class="detail-label">材质</span><span class="detail-value">${escapeHtml(item.material)}</span></div>
        <div class="detail-row"><span class="detail-label">透光度</span><span class="detail-value">${item.translucency.toFixed(2)}</span></div>
        <div class="detail-row"><span class="detail-label">细度</span><span class="detail-value">${item.fineness.toFixed(2)}</span></div>
        <div class="detail-row"><span class="detail-label">珠子颗数</span><span class="detail-value">${beadDisplay}</span></div>
        <div class="detail-row"><span class="detail-label">录入时间</span><span class="detail-value">${createdAt}</span></div>
        <div class="detail-row"><span class="detail-label">综合评分</span><span class="detail-value highlight">★ ${item.score.toFixed(2)} 分</span></div>
        <div class="detail-row"><span class="detail-label">评定等级</span><span><span class="grade-badge ${escapeHtml(item.grade)}">${escapeHtml(item.grade)}</span></span></div>
    `;
}

function downloadCard(id) {
    const link = document.createElement('a');
    link.href = `${API_BASE}/bracelets/${id}/card`;
    link.download = '';
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
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
