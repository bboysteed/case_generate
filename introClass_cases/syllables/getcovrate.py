import os
import re
import subprocess
import sys
from xml.dom.minidom import parse

sys.path.append(r"/home/steed/Desktop/session_work/case_generate/introClass_cases")
from utils.pycui import *

"""
    @module_name      全局路径设定
"""
module_name = "syllables"
module_base_path = f"/home/steed/Desktop/session_work/IntroClass/{module_name}"
module_testFile_path = os.path.join(module_base_path, "tests")
makeFile_path = os.path.join(module_base_path, "Makefile")
bench_execu_file_path = os.path.join(module_testFile_path, module_name)
cov_xml_path = os.path.join(module_testFile_path, "cov.xml")
cov_json_path = os.path.join(module_testFile_path, "cov.json")
"""
    @color      颜色输出模块
"""
color = pycui()


def check_syllables(test_output, truth_output,):
    answer_rgx = re.compile('syllables is (\d+)', re.IGNORECASE)
    truth_answer = answer_rgx.search(truth_output).group(1)
    right_answer = re.compile('syllables is {}'.format(truth_answer),
                              re.IGNORECASE)
    other_answers = re.compile('syllables is [^{}]'.format(
        truth_answer))
    right_answer_p = right_answer.search(test_output) is not None
    other_answers_p = other_answers.search(test_output) is None
    return right_answer_p and other_answers_p


def make_bench_file_clean():
    rewrite_makefile(module_testFile_path)
    color.success("清除编译痕迹...")
    os.system("make -C {} clean".format(module_testFile_path))


def make_test_file(test_file_path):
    color.success("编译被测可执行文件...")
    os.system("make -C {} all".format(test_file_path))


def rewrite_makefile(test_file_path):
    color.success("重写MakeFile...")
    makeFile = open(os.path.join(test_file_path, "Makefile"), "w")
    makeFile.write(
        f"CFLAGS = -lm -ftest-coverage -fprofile-arcs -fPIC\nall: {module_name}\nclean:\n\trm -f {module_name} {module_name}.gcda\n\trm -f {module_name}.gcno cov.xml")
    makeFile.close()


def make_clean_test_file(test_file_path):
    color.success("清除被测文件的编译痕迹...")
    os.system("make -C {} clean".format(test_file_path))


def make_bench_file():
    color.success("编译标准可执行文件...")
    os.system("make -C {} all".format(module_testFile_path))


def run_bench_file(inp_b_data):
    ret = subprocess.check_output(bench_execu_file_path, stdin=inp_b_data, timeout=2).decode()
    inp_b_data.seek(0)
    # color.info("bench file running result is:{}".format(ret))


def run_genprog_tests(test_exec_file_path, inp_data, inp_data_str):
    refer_res = subprocess.check_output(bench_execu_file_path, stdin=inp_data, timeout=3, shell=True).decode()
    inp_data.seek(0)  # 这真的是让人咂舌
    test_res = subprocess.check_output(test_exec_file_path, stdin=inp_data, timeout=3, shell=True).decode()
    color.info("refer_res:{}\n test_res:{}\n".format(refer_res, test_res))
    test_passed = check_syllables(test_output=test_res, truth_output=refer_res)
    color.warning(test_passed)
    color.info("Test {}.".format("passed" if test_passed else "failed"))
    return True if test_passed else False


def gcovr_save_xml():
    os.system("gcovr -r {} --xml-pretty -o {}/cov.xml".format(module_testFile_path, module_testFile_path))


def gcovr_save_test_xml(test_file_path):
    os.system("gcovr -r {} --xml-pretty -o {}/test_cov.xml".format(test_file_path, test_file_path))


def gcovr_save_json():
    os.system("gcovr -r {} --json-pretty -o {}/cov.json".format(module_testFile_path, module_testFile_path))


def parse_xml_and_get_rate():
    rootNode = parse(cov_xml_path).documentElement
    line_rate = rootNode.getAttribute("line-rate")  # 行覆盖率
    branch_rate = rootNode.getAttribute("branch-rate")  # 分支覆盖率
    lines_covered = rootNode.getAttribute("lines-covered")  # 覆盖行数
    branches_covered = rootNode.getAttribute("branches-covered")  # 覆盖分支数
    color.success("line_rate: {}".format(line_rate))
    return line_rate


if __name__ == '__main__':
    pass
