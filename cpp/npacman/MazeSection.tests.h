#pragma once
#include "Funcoes.h"
#include "MazeSection.h"

namespace my {

	namespace tests {

		namespace mzsection {

			void test_mzSection() {

				MazeSection s1(1,1);
				MazeSection s2(2, 1);
				MazeSection s3(1, 1);

				assertTrue(s1 == s3, "s1 == s3");
				assertTrue(s1 != s2, "s1 != s2");
				assertTrue(!(s1 == s2), "!(s1 == s2)");
			}
		}
	}
}