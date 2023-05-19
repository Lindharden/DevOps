"""
To run this test with a visible browser, the following dependencies have to be setup:

  * `pip install selenium`
  * `pip install pymongo`
  * `pip install pytest`
  * `wget https://github.com/mozilla/geckodriver/releases/download/v0.32.0/geckodriver-v0.32.0-linux64.tar.gz`
  * `tar xzvf geckodriver-v0.32.0-linux64.tar.gz`
  * After extraction, the downloaded artifact can be removed: `rm geckodriver-v0.32.0-linux64.tar.gz`

The application that it tests is the version of _ITU-MiniTwit_ that you got to know during the exercises on Docker:
https://github.com/itu-devops/flask-minitwit-mongodb/tree/Containerize (*OBS*: branch Containerize)

```bash
$ git clone https://github.com/HelgeCPH/flask-minitwit-mongodb.git
$ cd flask-minitwit-mongodb
$ git switch Containerize
```

After editing the `docker-compose.yml` file file where you replace `youruser` with your respective username, the
application can be started with `docker-compose up`.

Now, the test itself can be executed via: `pytest test_itu_minitwit_ui.py`.
"""

import pytest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.webdriver.firefox.service import Service
from selenium.webdriver.firefox.options import Options
from test_sim_compliance import wait_for_port
import os
import subprocess
import sqlite3

GUI_URL = "http://localhost:8080/register"
DATABASE = "itu-minitwit.db"

con = sqlite3.connect(DATABASE)

@pytest.fixture(scope='session', autouse=True)
def start_service():
    proc = subprocess.Popen(['go', 'run', 'minitwit.go'],
                            stdout=subprocess.DEVNULL,
                            stderr=subprocess.STDOUT,
                            shell=False)
    wait_for_port(8080, timeout=20)
    yield
    os.system("kill -9 $(lsof -t -i:8080)")

def _register_user_via_gui(driver, data):
    driver.get(GUI_URL)

    wait = WebDriverWait(driver, 5)
    buttons = wait.until(EC.presence_of_all_elements_located((By.CLASS_NAME, "actions")))
    input_fields = driver.find_elements(By.TAG_NAME, "input")

    for idx, str_content in enumerate(data):
        input_fields[idx].send_keys(str_content)
    input_fields[4].click()

    wait = WebDriverWait(driver, 5)
    flashes = wait.until(EC.presence_of_all_elements_located((By.TAG_NAME, "h2")))
    return flashes


def _get_user_by_name(name):

    try:
        cur = con.cursor()
        cur.execute("select * from users where username =?", (name,))
        rows = cur.fetchall()
        return rows[0]
    except Exception as e:
        return None


def test_register_user_via_gui():
    """
    This is a UI test. It only interacts with the UI that is rendered in the browser and checks that visual
    responses that users observe are displayed.
    """
    firefox_options = Options()
    firefox_options.add_argument("--headless")
    # firefox_options = None
    with webdriver.Firefox(service=Service("./test/geckodriver"), options=firefox_options) as driver:
        generated_msg = _register_user_via_gui(driver, ["Me", "me@some.where", "secure123", "secure123"])[0].text
        expected_msg = "Sign In"
        assert generated_msg == expected_msg



def test_register_user_via_gui_and_check_db_entry():
    """
    This is an end-to-end test. Before registering a user via the UI, it checks that no such user exists in the
    database yet. After registering a user, it checks that the respective user appears in the database.
    """
    firefox_options = Options()
    firefox_options.add_argument("--headless")
    # firefox_options = None
    with webdriver.Firefox(service=Service("./test/geckodriver"), options=firefox_options) as driver:

        assert _get_user_by_name("Me2") == None

        generated_msg = _register_user_via_gui(driver, ["Me2", "me@some.where", "secure123", "secure123"])[0].text
        expected_msg = "Sign In"
        assert generated_msg == expected_msg
   
        assert _get_user_by_name("Me2")[4] == "Me2"


