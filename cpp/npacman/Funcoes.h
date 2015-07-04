#pragma once
#include <iostream>

namespace my {

	namespace tests {

		//Apena imprime a mensagem informada
		void testCase(const char* name) {
			std::cout << name << std::endl;
		}

		//Verifica se uma certa condição é verdadeira
		void assertTrue(bool v, const char* error_msg) {
			if (!v) {
				std::cout << "[ERROR] " << error_msg << std::endl;
				return;
			}

			std::cout << "[OK   ] " << error_msg << std::endl;
		}
	}
}