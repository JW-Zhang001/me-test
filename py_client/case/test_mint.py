# -*- coding: utf-8 -*-
import json

import httpx
import yaml

with open("../config/config.yml", "r") as f:
    config = yaml.safe_load(f)


class TestMint(object):
    http_client = httpx.Client()

    with open('../../tools/open_api.json', "r") as f:
        open_api = json.load(f)

    def test_mint(self):
        get_account = self.open_api[0]
        get_account['url'] = config["api_url"] + get_account['url']
        resp = self.http_client.request(**get_account).json()['accounts']
        mint_info = [i for i in resp if i.get("name") == "mint"]

        mint_addr = mint_info[0]['base_account']['address']

        block_num_data = {"url": config.base_url + "/cosmos/base/tendermint/v1beta1/blocks/latest", "method": "GET"}
        block_resp = self.http_client.request(**block_num_data).json()
        block_height = block_resp['block']['header']['height']

        get_balances = {"url": config.base_url + f"/cosmos/bank/v1beta1/balances/{mint_addr}",
                        "method": "GET", }
        resp = self.http_client.request(**get_balances).json()

        result = int(block_height) * 792 * 10 ** 6
        assert int(resp['balances'][0]['amount']) == result
        pass
