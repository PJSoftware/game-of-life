import unittest

from life import world

class MyTest(unittest.TestCase):
    myWorld = world.World()

    def test_gridReference(self):
        self.assertEqual(self.myWorld._gridReference(1, 1), "1|1")
        self.assertEqual(self.myWorld._gridReference(128, 64), "128|64")

    def test_wrapX(self):
        self.assertEqual(self.myWorld._gridReference(0, 1), "128|1")
        self.assertEqual(self.myWorld._gridReference(129, 1), "1|1")

    def test_wrapY(self):
        self.assertEqual(self.myWorld._gridReference(1, 0), "1|64")
        self.assertEqual(self.myWorld._gridReference(1, 65), "1|1")

if __name__ == "__main__":
    unittest.main()
