# Tool: run_tests

Run the grading harness against the built server.

- name: `run_tests`
- command: `python -m bob.harness.runner --server 127.0.0.1:<port> --seed-set working`
- returns: pass rate per command group and overall, plus failing sequence seeds.
