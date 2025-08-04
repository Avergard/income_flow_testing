import pytest
import requests

BASE_URL = "http://localhost:8080/good"

@pytest.fixture
def sample_good():
    return {
        "id": 12,
        "name": f"Test Product",
        "description": "Test description",
        "count": 1,
        "weight": 1
    }

def test_get_all_goods():
    response = requests.post(f"{BASE_URL}/get_all")
    assert response.status_code == 200
    goods = response.json()
    assert isinstance(goods, list)
    assert len(goods) >= 1

def test_get_one_good():
    all_goods = requests.post(f"{BASE_URL}/get_all").json()
    test_id = all_goods[0]['id']

    response = requests.post(f"{BASE_URL}/get", json={"id": test_id})
    assert response.status_code == 200
    good = response.json()
    assert good['id'] == test_id

def test_create_good(sample_good):
    response = requests.post(f"{BASE_URL}/create", json=sample_good)
    assert response.status_code == 200
    created_good = response.json()

    assert created_good['name'] == sample_good['name']
    assert created_good['id'] == sample_good['id']

    all_goods = requests.post(f"{BASE_URL}/get_all").json()
    assert any(g['id'] == sample_good['id'] for g in all_goods)

    delete_response = requests.post(f"{BASE_URL}/delete", json={"id": sample_good["id"]})
    assert delete_response.status_code == 200

def test_update_good(sample_good):
    create_response = requests.post(f"{BASE_URL}/create", json=sample_good)
    assert create_response.status_code == 200

    updated_data = {
        "id": sample_good["id"],
        "name": "Updated Name",
        "description": "Updated description",
        "count": 99,
        "weight": 9.9
    }

    response = requests.post(f"{BASE_URL}/update", json=updated_data)
    assert response.status_code == 200
    updated_good = response.json()
    assert updated_good['name'] == "Updated Name"
    assert updated_good['count'] == 99
    assert updated_good['weight'] == 9.9

    requests.post(f"{BASE_URL}/delete", json={"id": sample_good["id"]})

def test_delete_good(sample_good):
    create_response = requests.post(f"{BASE_URL}/create", json=sample_good)
    assert create_response.status_code == 200

    delete_response = requests.post(f"{BASE_URL}/delete", json={"id": sample_good["id"]})
    assert delete_response.status_code == 200

    get_response = requests.post(f"{BASE_URL}/get", json={"id": sample_good["id"]})
    assert get_response.status_code == 404

def test_get_nonexistent_good():
    response = requests.post(f"{BASE_URL}/get", json={"id": 99999})
    assert response.status_code == 404
    error_data = response.json()
    assert error_data['code'] == "there_is_no_good"

def test_create_invalid_good():
    invalid_good = {
        "id": "not_an_integer",
        "name": "Test",
        "count": 1
    }

    response = requests.post(f"{BASE_URL}/create", json=invalid_good)
    assert response.status_code == 422
    error_data = response.json()
    assert error_data['code'] == "marshal_error"

def test_create_duplicate_id():
    good1 = {
        "id": 100,
        "name": "Test1",
        "description": "Test1",
        "count": 1,
        "weight": 1.0
    }
    requests.post(f"{BASE_URL}/create", json=good1)

    good2 = {
        "id": 100,
        "name": "Test2",
        "description": "Test2",
        "count": 2,
        "weight": 2.0
    }
    response = requests.post(f"{BASE_URL}/create", json=good2)
    assert response.status_code == 400
    error_data = response.json()
    assert error_data['code'] == "User with this id already exists"

def test_get_one_good_empty():
    response = requests.post(f"{BASE_URL}/get", json={"id": None})
    assert response.status_code == 422
    error_data = response.json()
    assert error_data['code'] == "marshal_error"