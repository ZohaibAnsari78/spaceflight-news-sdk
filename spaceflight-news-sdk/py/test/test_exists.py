# SpaceflightNews SDK exists test

import pytest
from spaceflightnews_sdk import SpaceflightNewsSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = SpaceflightNewsSDK.test(None, None)
        assert testsdk is not None
