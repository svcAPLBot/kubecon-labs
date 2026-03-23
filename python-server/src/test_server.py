import unittest
import os
from server import response_message

class TestHelloHandler(unittest.TestCase):
    def test_response_uses_env_values(self):
        os.environ["FEELING"] = "curious"
        os.environ["AWESOMNESS"] = "snacks"

        response = response_message()

        self.assertEqual(
            response,
            'Howdy! I feeeeel curious today. I will tell you my secret of awesomeness: "snacks".',
        )
    
    def test_default_message(self):
        os.environ.pop("FEELING", None)
        os.environ.pop("AWESOMNESS", None)

        self.assertEqual(
            response_message(),
            'Howdy! I feeeeel $FEELING today. I will tell you my secret of awesomeness: "$AWESOMNESS".',
        )

if __name__ == "__main__":
    unittest.main()
