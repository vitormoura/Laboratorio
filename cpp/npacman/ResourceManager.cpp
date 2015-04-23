#include "ResourceManager.h"
#include <fstream>
#include <streambuf>


namespace my {
	
	ResourceManager::ResourceManager()
	{
	}
	
	ResourceManager::~ResourceManager()
	{
	}

	std::unique_ptr<std::string> ResourceManager::getFileContents(const char* filePath) {
		
		std::ifstream t(filePath);
		std::string* result = new std::string();

		//Verificando tamanho do arquivo para alocar previamente o espaço necessário em nossa string
		t.seekg(0, std::ios::end);
		result->reserve(t.tellg());
		t.seekg(0, std::ios::beg);

		result->assign(std::istreambuf_iterator<char>(t), std::istreambuf_iterator<char>());

		return std::unique_ptr<std::string>(result);
	}
}